package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Show the values selected for the HTTP request",
	Long: `Show the values (Host, Port, Path, Method and Body) for the HTTP request,
when none is given it show all the values.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			showAllValues()
		}
	},
}

var checkHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Show the host setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		printTitledValue("Host:", viper.GetViper().GetString("host"))
	},
}

var checkPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show the path setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		printTitledValue("Path:", viper.GetViper().GetString("path"))
	},
}

var checkPortCmd = &cobra.Command{
	Use:   "port",
	Short: "Show the port setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		portStr := strconv.Itoa(viper.GetViper().GetInt("port"))
		printTitledValue("Port:", portStr)
	},
}

var checkMethodCmd = &cobra.Command{
	Use:   "method",
	Short: "Show the method setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		printTitledValue("Method:", viper.GetViper().GetString("method"))
	},
}

var checkBodyCmd = &cobra.Command{
	Use:   "body",
	Short: "Show the body setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		printTitledValue("Body:", viper.GetViper().GetString("body"))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.AddCommand(checkHostCmd)
	checkCmd.AddCommand(checkPathCmd)
	checkCmd.AddCommand(checkPortCmd)
	checkCmd.AddCommand(checkMethodCmd)
	checkCmd.AddCommand(checkBodyCmd)
}
