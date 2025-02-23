/*
Copyright © 2025 Axel Aparicio
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

var rootCmd = &cobra.Command{
	Use:   "gato",
	Short: "Gato is an HTTP request client CLI similar to cURL",
	Long:  `Gato is an HTTP request client CLI similar to cURL`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigFile(cfgFile)

	_ = viper.ReadInConfig()
}
