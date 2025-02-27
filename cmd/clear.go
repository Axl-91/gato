package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all values and restore them to defaults",
	Long:  `Clear all values and restore them to the default ones`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(rootCmd.OutOrStderr(), "Invalid amount of parameter, 1 or none expected")
			return
		}
		parameter := extractParameter(args)
		clearValues(parameter)
	},
}

func clearValues(parameter string) {
	switch parameter {
	case "host":
		viper.GetViper().Set("host", "http://127.0.0.1")
		printTitledValue("host cleared to:", "http://127.0.0.1")
	case "port":
		viper.GetViper().Set("port", 8000)
		printTitledValue("port cleared to:", "8000")
	case "path":
		viper.GetViper().Set("path", "")
		printTitledValue("path cleared to:", "")
	case "method":
		viper.GetViper().Set("method", "GET")
		printTitledValue("method cleared to:", "GET")
	case "body":
		viper.GetViper().Set("body", "")
		printTitledValue("body cleared to:", "")
	case "":
		clearAllValues()
	}
	_ = viper.GetViper().WriteConfig()
}

func clearAllValues() {
	viper.GetViper().Set("host", "http://127.0.0.1")
	viper.GetViper().Set("path", "")
	viper.GetViper().Set("port", 8000)
	viper.GetViper().Set("method", "GET")
	viper.GetViper().Set("body", "")
	fmt.Fprintln(rootCmd.OutOrStdout(), titleStyle.Render("All values cleared"))
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
