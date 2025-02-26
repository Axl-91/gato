package cmd

import (
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
		portStr := strconv.Itoa(viper.GetViper().GetInt("port"))

		printValue("Host:", viper.GetViper().GetString("host"))
		printValue("Port:", portStr)
		printValue("Path:", viper.GetViper().GetString("path"))
		printValue("Method:", viper.GetViper().GetString("method"))
		printValue("Body:", viper.GetViper().GetString("body"))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
