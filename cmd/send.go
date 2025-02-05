/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var client = http.Client{
	Timeout: time.Duration(10 * time.Second),
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send the request",
	Long: `Send the request with the specifications selected
		- Host (default: http://localhost.com)
		- Dir (default: None)
		- Port (default: 8080)
		- Method (default: GET)
		- Body (default: None)
		`,
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetViper().GetString("host")
		port := viper.GetViper().GetString("port")
		dir := viper.GetViper().GetString("dir")

		url := host + ":" + port + "/" + dir
		method := viper.GetViper().GetString("method")

		switch method {
		case "GET":
			get_request(url)
		case "POST":
			post_request(url)
		}
	},
}

func get_request(url string) {
	fmt.Println("Sending Request to", url)
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal("Error on request:", err)
	}
	defer resp.Body.Close()

	statusCode := resp.StatusCode
	fmt.Println("STATUS CODE:", statusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error on the request body:", err)
	}

	fmt.Println("RESPONSE: ")
	fmt.Println(string(body))
}

func post_request(url string) {

}
func init() {
	rootCmd.AddCommand(sendCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
