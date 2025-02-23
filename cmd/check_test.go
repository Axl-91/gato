package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestCheckCommand(t *testing.T) {
	defaultValues := []string{"http://127.0.0.1", "", "8000", "GET", ""}
	newValues := []string{"http://google.com", "/api/storage", "4000", "POST", "body.json"}

	// Check that the values are the default ones.
	setValues(defaultValues)
	checkValues(t, defaultValues)

	// I'll change the values and check that they are the same
	setValues(newValues)
	checkValues(t, newValues)

	// I'll restore the values and check that they are the default ones
	setValues(defaultValues)
	checkValues(t, defaultValues)
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
		if !strings.Contains(output, value) {
			if value == "" {
				value = "None"
			}
			t.Errorf("Expected output to contain %q but got %q", value, output)
		}
	}
}

func setValues(values []string) {
	viper.GetViper().Set("host", values[0])
	viper.GetViper().Set("path", values[1])
	viper.GetViper().Set("port", values[2])
	viper.GetViper().Set("method", values[3])
	viper.GetViper().Set("body", values[4])

	_ = viper.GetViper().WriteConfig()
}
