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
		port_str := strconv.Itoa(viper.GetViper().GetInt("port"))

		printValue("Host:", viper.GetViper().GetString("host"))
		printValue("Port:", port_str)
		printValue("Path:", viper.GetViper().GetString("path"))
		printValue("Method:", viper.GetViper().GetString("method"))
		printValue("Body:", viper.GetViper().GetString("body"))
	},
}

func printValue(name string, value string) {
	output := rootCmd.OutOrStdout()
	name_formatted := title_style.Render(name)

	fmt.Fprintf(output, "%s %s\n", name_formatted, valueOrNone(value))
}

func valueOrNone(value string) string {
	if value == "" {
		return "None"
	}
	return value
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
