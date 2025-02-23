package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Show all the values for the request",
	Long:  `Show all the values (Host, Port, Path, Method and Body) for the request`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(rootCmd.OutOrStdout(), "%s %s \n", title_style.Render("Host:"), viper.GetViper().GetString("host"))
		fmt.Fprintf(rootCmd.OutOrStdout(), "%s %d \n", title_style.Render("Port:"), viper.GetViper().GetInt("port"))
		fmt.Fprintf(rootCmd.OutOrStdout(), "%s %s \n", title_style.Render("Path:"), value_or_none(viper.GetViper().GetString("path")))
		fmt.Fprintf(rootCmd.OutOrStdout(), "%s %s \n", title_style.Render("Method:"), viper.GetViper().GetString("method"))
		fmt.Fprintf(rootCmd.OutOrStdout(), "%s %s \n", title_style.Render("Body:"), value_or_none(viper.GetViper().GetString("body")))
	},
}

func value_or_none(value string) string {
	if value == "" {
		return "None"
	}
	return value
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
