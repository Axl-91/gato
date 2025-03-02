package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	if err := os.Chdir(".."); err != nil {
		fmt.Println("Failed to change directory:", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}

func setValues(host string, path string, port int, method string, body string) {
	viper.GetViper().Set("host", host)
	viper.GetViper().Set("path", path)
	viper.GetViper().Set("port", port)
	viper.GetViper().Set("method", method)
	viper.GetViper().Set("body", body)

	_ = viper.GetViper().WriteConfig()
}

func checkValues(t *testing.T, checkValues []string) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"check"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error executing check command: %v", err)
	}

	output := buf.String()

	for _, value := range checkValues {
		if value == "" {
			value = "None"
		}
		if !strings.Contains(output, value) {
			t.Errorf("Expected output to contain %q but got %q", value, output)
		}
	}
}
