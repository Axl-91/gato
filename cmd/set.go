/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var host string
var dir string
var port int16
var method string
var body string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Command used to set different values needed in requests",
	Long:  `This command will be used to set all different values needed`,
	Run: func(cmd *cobra.Command, args []string) {
		if host != "" {
			if strings.Contains(host, "localhost") {
				host = "http://127.0.0.1"
			}
			if !strings.Contains(host, "http") {
				host = "http://" + host
			}
			viper.GetViper().Set("host", host)
			fmt.Println(title_style.Render("Host changed to:"), host)
		}
		if dir != "" {
			viper.GetViper().Set("dir", dir)
			fmt.Println(title_style.Render("Dir changed to:"), dir)
		}
		if port >= 0 {
			viper.GetViper().Set("port", port)
			fmt.Println(title_style.Render("Port changed to:"), port)
		}
		if method != "" {
			valid_methods := []string{"GET", "POST"}
			if slices.Contains(valid_methods, method) {
				viper.GetViper().Set("method", method)
				fmt.Println(title_style.Render("Method changed to:"), method)
			} else {
				fmt.Println(error_style.Render("Invalid Method Selected"))
			}
		}
		if body != "" {
			viper.GetViper().Set("body", body)
			fmt.Println(title_style.Render("Body changed to:"), body)
		}
		_ = viper.GetViper().WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	setCmd.Flags().StringVarP(&host, "host", "H", "", "Set host name")
	setCmd.Flags().StringVarP(&dir, "dir", "D", "", "Set host dir")
	setCmd.Flags().Int16VarP(&port, "port", "P", -1, "Set port number")
	setCmd.Flags().StringVarP(&method, "method", "M", "", "Set request method")
	setCmd.Flags().StringVarP(&body, "body", "B", "", "Set body (json)")
	setCmd.MarkFlagsOneRequired("host", "port", "method", "body")
}
