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
		fmt.Println(title_style.Render("Dir:"), viper.GetViper().GetString("dir"))
		fmt.Println(title_style.Render("Method:"), viper.GetViper().GetString("method"))
		fmt.Println(title_style.Render("Body:"), viper.GetViper().GetString("body"))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
