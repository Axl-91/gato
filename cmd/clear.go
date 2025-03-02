package cmd

import (
	"fmt"
	"strconv"

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
			fmt.Fprintln(rootCmd.OutOrStderr(), "Invalid amount of parameters, 1 or none expected")
			return
		}
		parameter := extractParameter(args)
		clearValues(parameter)
	},
}

func clearValues(parameter string) {
	switch parameter {
	case "host":
		viper.GetViper().Set("host", defaultValues.Host)
		printTitledValue("Host cleared to:", defaultValues.Host)
	case "port":
		viper.GetViper().Set("port", defaultValues.Port)
		portStr := strconv.Itoa(int(defaultValues.Port))
		printTitledValue("Port cleared to:", portStr)
	case "path":
		viper.GetViper().Set("path", defaultValues.Path)
		printTitledValue("Path cleared to:", defaultValues.Path)
	case "method":
		viper.GetViper().Set("method", defaultValues.Method)
		printTitledValue("Method cleared to:", defaultValues.Method)
	case "body":
		viper.GetViper().Set("body", defaultValues.Body)
		printTitledValue("Body cleared to:", defaultValues.Body)
	case "":
		clearAllValues()
	}
	_ = viper.GetViper().WriteConfig()
}

func clearAllValues() {
	viper.GetViper().Set("host", defaultValues.Host)
	viper.GetViper().Set("path", defaultValues.Path)
	viper.GetViper().Set("port", defaultValues.Port)
	viper.GetViper().Set("method", defaultValues.Method)
	viper.GetViper().Set("body", defaultValues.Body)
	fmt.Fprintln(rootCmd.OutOrStdout(), titleStyle.Render("All values cleared"))
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
