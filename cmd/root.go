/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string = "gato.yaml"

var title_style = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("#36c7aa"))

var var_style = lipgloss.NewStyle().Italic(true)

var error_style = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("9"))

var status_style = lipgloss.NewStyle().Bold(true).
	Background(lipgloss.Color("#3bcc06"))

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gato",
	Short: "Gato CLI is a request application similar to cURL",
	Long:  `Gato CLI is a request application similar to cURL`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)

	_ = viper.ReadInConfig()
}
