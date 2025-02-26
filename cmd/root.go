package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "gato.yaml"

// Style for titles
var titleStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#36c7aa"))

// Style for error messages
var errorStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("9"))

// Style to show status (default: background green)
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
	viper.SetConfigFile(cfgFile)

	_ = viper.ReadInConfig()
}

func printValue(name string, value string) {
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
