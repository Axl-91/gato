/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/maps"
)

var client = http.Client{
	Timeout: time.Duration(5 * time.Second),
}

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send the request",
	Long: `Send the request with the specifications selected
- Host (default: http://localhost.com)
- Path (default: None)
- Port (default: 8080)
- Method (default: GET)
- Body (default: None)
	`,
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetViper().GetString("host")
		path := viper.GetViper().GetString("path")
		port := viper.GetViper().GetString("port")
		method := viper.GetViper().GetString("method")
		body := viper.GetViper().GetString("body")

		url := host + ":" + port + "/" + path
		fmt.Println(title_style.Render("Sending Request... "))
		fmt.Print(title_style.Render("Host: "))
		fmt.Print(var_style.Render(url))
		fmt.Println()

		switch method {
		case "GET":
			get_request(url)
		case "POST":
			post_request(url, body)
		}
	},
}

func get_request(url string) {
	fmt.Print(title_style.Render("Method: "))
	fmt.Print(var_style.Render("GET"))
	fmt.Println()

	resp, err := client.Get(url)
	show_error_message(err)

	defer resp.Body.Close()

	show_parsed_response(resp)
}

func post_request(url_string string, body_json string) {
	fmt.Print(title_style.Render("Method: "))
	fmt.Print(var_style.Render("GET"))
	fmt.Println()

	// Read the json file to use it as a body for the POST request
	body, err := os.ReadFile(body_json)
	if err != nil {
		log.Fatal("Error trying to read json body:", err)
	}
	resp, err := client.Post(url_string, "application/json", bytes.NewBuffer(body))

	show_error_message(err)

	defer resp.Body.Close()

	show_parsed_response(resp)
}

func show_error_message(err error) {
	if err != nil {
		fmt.Println("Error on request:")
		if strings.Contains(err.Error(), "connection refused") {
			fmt.Println("Connection Refused")
		} else {
			fmt.Println(err)
		}
		return
	}
}

func show_parsed_response(response *http.Response) {
	show_status_code(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error on the request body: ", err)
	}

	show_table_response(body)
}

func show_status_code(status_code int) {
	if status_code >= 400 {
		status_style = status_style.Background(lipgloss.Color("9"))
	}

	fmt.Println(
		title_style.Render("STATUS CODE:"),
		status_style.Render(strconv.Itoa(status_code)),
	)
}

func show_table_response(body []byte) {
	fmt.Println(title_style.Render("RESPONSE: "))

	var parsed_list_body []map[string]string
	var parsed_map_body map[string]string
	_ = json.Unmarshal(body, &parsed_list_body)
	_ = json.Unmarshal(body, &parsed_map_body)

	parsed_list_body = append(parsed_list_body, parsed_map_body)

	var list_keys []string
	var list_values [][]string

	if len(parsed_list_body) > 0 {
		list_keys = maps.Keys(parsed_list_body[0])
	}

	for _, body_map := range parsed_list_body {
		var list_map []string
		for _, key := range list_keys {
			list_map = append(list_map, body_map[key])
		}
		list_values = append(list_values, list_map)
	}

	table := create_table(list_keys, list_values)

	fmt.Println(table)
}

func create_table(list_keys []string, list_values [][]string) *table.Table {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == -1 {
				return title_style
			}
			return lipgloss.NewStyle()
		}).
		Headers(list_keys...).
		Rows(list_values...)
	return t
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
