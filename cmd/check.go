/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
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
	Long:  `Show all the values (Host, Port, Dir, Method and Body) for the request`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(title_style.Render("Host:"), viper.GetViper().GetString("host"))
		fmt.Println(title_style.Render("Port:"), viper.GetViper().GetInt("port"))
		fmt.Println(title_style.Render("Dir:"), value_or_none(viper.GetViper().GetString("dir")))
		fmt.Println(title_style.Render("Method:"), viper.GetViper().GetString("method"))
		fmt.Println(title_style.Render("Body:"), value_or_none(viper.GetViper().GetString("body")))
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
