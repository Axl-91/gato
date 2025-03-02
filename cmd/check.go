package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Show all the values for the request",
	Long:  `Show all the values (Host, Port, Path, Method and Body) for the request`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Fprintln(rootCmd.OutOrStderr(), "Invalid amount of parameters, 1 or none expected")
			return
		}

		parameter := extractParameter(args)
		showValues(parameter)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
