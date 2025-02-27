package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var host string
var path string
var port int16
var method string
var body string

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
			printValue("Host changed to:", host)
		}
		if path != "" {
			viper.GetViper().Set("path", path)
			printValue("Path changed to:", path)
		}
		if port >= 0 {
			viper.GetViper().Set("port", port)
			portStr := strconv.Itoa(int(port))
			printValue("Port changed to:", portStr)
		}
		if method != "" {
			validMethods := []string{"GET", "POST"}
			if slices.Contains(validMethods, method) {
				viper.GetViper().Set("method", method)
				printValue("Method changed to:", method)
			} else {
				errorMsg := fmt.Sprintf("Invalid method selected: expected [%s], got %s", strings.Join(validMethods, ", "), method)
				errorFormatted := errorStyle.Render(errorMsg)
				fmt.Fprintf(rootCmd.ErrOrStderr(), "%s\n", errorFormatted)
			}
		}
		if body != "" {
			viper.GetViper().Set("body", body)
			printValue("Body changed to:", body)
		}
		_ = viper.GetViper().WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&host, "host", "H", "", "Set host name")
	setCmd.Flags().StringVarP(&path, "path", "D", "", "Set host path")
	setCmd.Flags().Int16VarP(&port, "port", "P", -1, "Set port number")
	setCmd.Flags().StringVarP(&method, "method", "M", "", "Set request method")
	setCmd.Flags().StringVarP(&body, "body", "B", "", "Set body (json)")
	setCmd.MarkFlagsOneRequired("host", "path", "port", "method", "body")
}
