package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "gato.yaml"

type defaultJson struct {
	Host   string `json:"host"`
	Path   string `json:"path"`
	Port   int32  `json:"port"`
	Method string `json:"method"`
	Body   string `json:"body"`
}

var defaultValues = defaultJson{}

// Style for titles
var titleStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#36c7aa"))

// Style for success messages
var successStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#3bcc06"))

// Style for error messages
var errorStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("9"))

var statusStyle = lipgloss.NewStyle().Bold(true)

var rootCmd = &cobra.Command{
	Use:   "gato",
	Short: "Gato is an HTTP request client CLI similar to cURL",
	Long:  `Gato is an HTTP request client CLI similar to cURL`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	err := parseJsonDefault()
	if err != nil {
		os.Exit(1)
	}
	viper.SetConfigFile(cfgFile)

	_ = viper.ReadInConfig()
}

// Functions with multiple calls on different files will be here

func parseJsonDefault() error {
	file, err := os.Open("default.json")
	if err != nil {
		errorMsg := fmt.Sprintf(errorStyle.Render("Error opening default JSON file: %v"), err)
		fmt.Fprintln(rootCmd.OutOrStderr(), errorMsg)
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&defaultValues)
	if err != nil {
		errorMsg := fmt.Sprintf(errorStyle.Render("Error decoding JSON file: %v"), err)
		fmt.Fprintln(rootCmd.OutOrStderr(), errorMsg)
		return err
	}
	return nil
}

func printTitledValue(name string, value string) {
	output := rootCmd.OutOrStdout()
	nameFormatted := titleStyle.Render(name)

	fmt.Fprintf(output, "%s %s\n", nameFormatted, valueOrNone(value))
}

func valueOrNone(value string) string {
	if value == "" {
		return "None"
	}
	return value
}

func setStatusCodeStyle(statusCode int) {
	if statusCode >= 400 {
		statusStyle = statusStyle.Background(lipgloss.Color("9"))
	} else {
		statusStyle = statusStyle.Background(lipgloss.Color("#3bcc06"))

	}
}

func getValueString(value interface{}) string {
	switch v := value.(type) {
	case bool:
		return fmt.Sprintf("%t", v)
	case int, int64, float64:
		return fmt.Sprintf("%v", v)
	default:
		return fmt.Sprintf("%s", v)
	}
}

func extractParameter(args []string) string {
	if len(args) == 0 {
		return ""
	} else {
		return strings.ToLower(args[0])
	}
}

func showValues(parameter string) {
	switch parameter {
	case "host":
		printTitledValue("Host:", viper.GetViper().GetString("host"))
	case "port":
		portStr := strconv.Itoa(viper.GetViper().GetInt("port"))
		printTitledValue("Port:", portStr)
	case "path":
		printTitledValue("Path:", viper.GetViper().GetString("path"))
	case "method":
		printTitledValue("Method:", viper.GetViper().GetString("method"))
	case "body":
		printTitledValue("Body:", viper.GetViper().GetString("body"))
	case "":
		showAllValues()
	}
}

func showAllValues() {
	printTitledValue("Host:", viper.GetViper().GetString("host"))
	portStr := strconv.Itoa(viper.GetViper().GetInt("port"))
	printTitledValue("Port:", portStr)
	printTitledValue("Path:", viper.GetViper().GetString("path"))
	printTitledValue("Method:", viper.GetViper().GetString("method"))
	printTitledValue("Body:", viper.GetViper().GetString("body"))
}
