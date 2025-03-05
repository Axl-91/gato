package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear all values and restore them to defaults",
	Long:  `Clear all values and restore them to the default ones`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			clearAllValues()
		}
	},
}

var clearHostCmd = &cobra.Command{
	Use:   "host",
	Short: "Clear the host setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("host", defaultValues.Host)
		printTitledValue("Host cleared to:", defaultValues.Host)
		_ = viper.GetViper().WriteConfig()
	},
}

var clearPathCmd = &cobra.Command{
	Use:   "path",
	Short: "Clear the path setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("path", defaultValues.Path)
		printTitledValue("Path cleared to:", defaultValues.Path)
		_ = viper.GetViper().WriteConfig()
	},
}

var clearPortCmd = &cobra.Command{
	Use:   "port",
	Short: "Clear the port setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("port", defaultValues.Port)
		portStr := strconv.Itoa(int(defaultValues.Port))
		printTitledValue("Port cleared to:", portStr)
		_ = viper.GetViper().WriteConfig()
	},
}

var clearMethodCmd = &cobra.Command{
	Use:   "method",
	Short: "Clear the method setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("method", defaultValues.Method)
		printTitledValue("Method cleared to:", defaultValues.Method)
		_ = viper.GetViper().WriteConfig()
	},
}

var clearBodyCmd = &cobra.Command{
	Use:   "body",
	Short: "Clear the body setted for the HTTP request",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		viper.GetViper().Set("body", defaultValues.Body)
		printTitledValue("Body cleared to:", defaultValues.Body)
		_ = viper.GetViper().WriteConfig()
	},
}

func clearAllValues() {
	viper.GetViper().Set("host", defaultValues.Host)
	viper.GetViper().Set("path", defaultValues.Path)
	viper.GetViper().Set("port", defaultValues.Port)
	viper.GetViper().Set("method", defaultValues.Method)
	viper.GetViper().Set("body", defaultValues.Body)
	fmt.Fprintln(rootCmd.OutOrStdout(), titleStyle.Render("All values cleared"))
}

func init() {
	rootCmd.AddCommand(clearCmd)
	clearCmd.AddCommand(clearHostCmd)
	clearCmd.AddCommand(clearPathCmd)
	clearCmd.AddCommand(clearPortCmd)
	clearCmd.AddCommand(clearMethodCmd)
	clearCmd.AddCommand(clearBodyCmd)
}
