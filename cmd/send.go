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
		printValue("URL:", url)

		switch method {
		case "GET":
			getRequest(url)
		case "POST":
			postRequest(url, body)
		}
	},
}

func getRequest(url string) {
	printValue("Method:", "GET")

	resp, err := client.Get(url)
	showErrorMsg(err)

	defer resp.Body.Close()

	showParsedResp(resp)
}

func postRequest(urlString string, bodyJson string) {
	printValue("Method:", "POST")

	// Read the json file to use it as a body for the POST request
	body, err := os.ReadFile(bodyJson)
	if err != nil {
		log.Fatal("Error trying to read json body:", err)
	}
	resp, err := client.Post(urlString, "application/json", bytes.NewBuffer(body))

	showErrorMsg(err)

	defer resp.Body.Close()

	showParsedResp(resp)
}

func showParsedResp(response *http.Response) {
	showStatusCode(response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error on the request body: ", err)
	}

	showTableResp(body)
}

func showStatusCode(statusCode int) {
	setStatusCodeStyle(statusCode)

	statusCodeStr := strconv.Itoa(statusCode)

	printValue("STATUS CODE:", statusStyle.Render(statusCodeStr))
}

func showTableResp(body []byte) {
	// listBody will be used to retrieve response if is a list
	var listBody []map[string]interface{}

	// mapBody will be used to retrieve response if is an unique response
	var mapBody map[string]interface{}

	_ = json.Unmarshal(body, &listBody)
	_ = json.Unmarshal(body, &mapBody)

	listBody = append(listBody, mapBody)

	var listKeys []string
	var listValues [][]string

	if len(listBody) > 0 {
		listKeys = maps.Keys(listBody[0])
	}

	for _, bodyMap := range listBody {
		if len(bodyMap) == 0 {
			break
		}
		var listMap []string
		for _, key := range listKeys {
			valueString := getValueString(bodyMap[key])
			listMap = append(listMap, valueString)
		}
		listValues = append(listValues, listMap)
	}

	table := createTable(listKeys, listValues)

	fmt.Println(table)
}

func createTable(listKeys []string, listValues [][]string) *table.Table {
	t := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("99"))).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == -1 {
				return titleStyle
			}
			return lipgloss.NewStyle()
		}).
		Headers(listKeys...).
		Rows(listValues...)
	return t
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
