package cmd

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var host string
var path string
var port int16
var method string
var body string

type YamlRequest struct {
	Host   string `yaml:"host"`
	Path   string `yaml:"path"`
	Port   int16  `yaml:"port"`
	Method string `yaml:"method"`
	Body   string `yaml:"body"`
}

var setCmd = &cobra.Command{
	Use:   "set -flag new_value",
	Short: "Command used to set different values needed in requests",
	Long:  `This command will be used to set all different values needed`,
	Run: func(cmd *cobra.Command, args []string) {
		setRequestValues(host, path, port, method, body)
		fmt.Fprintln(rootCmd.OutOrStderr(), successStyle.Render("Value/s changed succesfully:"))
		showAllValues()
	},
}

var yamlCmd = &cobra.Command{
	Use:   "yaml <file>",
	Short: "Set request values from a YAML file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]

		data, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading YAML file: %v\n", err)
			return
		}

		var yamlFile YamlRequest
		err = yaml.Unmarshal(data, &yamlFile)
		if err != nil {
			fmt.Printf("Error parsing YAML: %v\n", err)
			return
		}
		setRequestValues(yamlFile.Host, yamlFile.Path, yamlFile.Port, yamlFile.Method, yamlFile.Body)
		fmt.Fprintln(rootCmd.OutOrStderr(), successStyle.Render("Value/s changed succesfully:"))
		showAllValues()
	},
}

func setRequestValues(host string, path string, port int16, method string, body string) {
	if host != "" {
		if strings.Contains(host, "localhost") {
			host = "http://127.0.0.1"
		}
		if !strings.Contains(host, "http") {
			host = "http://" + host
		}
		viper.GetViper().Set("host", host)
	}
	if path != "" {
		viper.GetViper().Set("path", path)
	}
	if port >= 0 {
		viper.GetViper().Set("port", port)
	}
	if method != "" {
		validMethods := []string{"GET", "POST"}
		if slices.Contains(validMethods, method) {
			viper.GetViper().Set("method", method)
		} else {
			errorMsg := fmt.Sprintf("Invalid method selected: expected [%s], got %s", strings.Join(validMethods, ", "), method)
			errorFormatted := errorStyle.Render(errorMsg)
			fmt.Fprintf(rootCmd.ErrOrStderr(), "%s\n", errorFormatted)
		}
	}
	if body != "" {
		viper.GetViper().Set("body", body)
	}
	_ = viper.GetViper().WriteConfig()

}

func init() {
	setCmd.Flags().StringVarP(&host, "host", "H", "", "Set host name")
	setCmd.Flags().StringVarP(&path, "path", "D", "", "Set host path")
	setCmd.Flags().Int16VarP(&port, "port", "P", -1, "Set port number")
	setCmd.Flags().StringVarP(&method, "method", "M", "", "Set request method")
	setCmd.Flags().StringVarP(&body, "body", "B", "", "Set body (json)")
	setCmd.MarkFlagsOneRequired("host", "path", "port", "method", "body")

	rootCmd.AddCommand(setCmd)
	setCmd.AddCommand(yamlCmd)
}
