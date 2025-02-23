/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all values and restore them to defaults",
	Long:  `Clear all values and restore them to the default ones`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("host", "http://127.0.0.1")
		viper.GetViper().Set("path", "")
		viper.GetViper().Set("port", 8000)
		viper.GetViper().Set("method", "GET")
		viper.GetViper().Set("body", "")
		_ = viper.GetViper().WriteConfig()

		fmt.Println("All values cleared")
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
