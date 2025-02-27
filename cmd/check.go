package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Show all the values for the request",
	Long:  `Show all the values (Host, Port, Path, Method and Body) for the request`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(rootCmd.OutOrStderr(), "Invalid amount of parameter, 1 or none expected")
			return
		}

		parameter := extractParameter(args)
		showValues(parameter)
	},
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

func init() {
	rootCmd.AddCommand(checkCmd)
}
