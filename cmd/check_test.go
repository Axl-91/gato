package cmd

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestCheckCommand(t *testing.T) {
	testCheckWithInvalidArgs(t)
	testCheckShowsAllValues(t)

	testCheckIndividualValue(t, "host", "http://127.0.0.1", "http://google.com")
	testCheckIndividualValue(t, "path", "None", "api/storage")
	testCheckIndividualValue(t, "port", "8000", "4000")
	testCheckIndividualValue(t, "method", "GET", "POST")
	testCheckIndividualValue(t, "body", "None", "body.json")
}

// Check for invalid amount of args
func testCheckWithInvalidArgs(t *testing.T) {
	testCheckWithInvalidAmount(t)
	testCheckWithInvalidMethod(t)
}

func testCheckWithInvalidAmount(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"check", "port", "another"})

	_ = rootCmd.Execute()
	expectedMsg := "accepts 0 arg(s), received 1"

	if !strings.Contains(buf.String(), expectedMsg) {
		t.Errorf("Expected output to contain %q but got %q", expectedMsg, buf.String())
	}
}

func testCheckWithInvalidMethod(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"check", "invalid"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error executing check command: %v", err)
	}
	output := buf.String()

	if len(output) > 0 {
		t.Errorf("Expected empty response, instead got: %s", output)
	}
}

// Check with none args shows all values of request
func testCheckShowsAllValues(t *testing.T) {
	defaultValues := []string{"http://127.0.0.1", "", "8000", "GET", ""}
	newValues := []string{"http://google.com", "/api/storage", "4000", "POST", "body.json"}

	newPort, _ := strconv.Atoi(newValues[2])

	// Check that the values are the default ones.
	rootCmd.SetArgs([]string{"clear"})
	_ = rootCmd.Execute()
	checkValues(t, defaultValues)

	// Change the values and use check command to validate they were changed
	setValues(newValues[0], newValues[1], newPort, newValues[3], newValues[4])
	checkValues(t, newValues)

	// Restore default values with the clear commnad
	rootCmd.SetArgs([]string{"clear"})
	_ = rootCmd.Execute()

	// Now we check that the values are the default ones again
	checkValues(t, defaultValues)
}

func testCheckIndividualValue(t *testing.T, value string, defaultVal string, newVal string) {
	var buffer bytes.Buffer

	// Check the value is default
	rootCmd.SetOut(&buffer)
	rootCmd.SetArgs([]string{"check", value})
	_ = rootCmd.Execute()

	response := buffer.String()
	expectedResponse := fmt.Sprintf(": %s\n", defaultVal)

	if !strings.Contains(response, expectedResponse) {
		t.Errorf("Expected '%s', instead got: '%s'", response, expectedResponse)
	}

	// Clear buffer
	buffer.Truncate(0)

	// Change value to new one
	viper.GetViper().Set(value, newVal)
	_ = viper.GetViper().WriteConfig()

	// Check the value is changed
	rootCmd.SetArgs([]string{"check", value})
	_ = rootCmd.Execute()

	response = buffer.String()
	expectedResponse = fmt.Sprintf(": %s\n", newVal)

	if !strings.Contains(response, expectedResponse) {
		t.Errorf("Expected '%s', instead got: '%s'", response, expectedResponse)
	}

	// Restore default values with the clear commnad
	rootCmd.SetArgs([]string{"clear", value})
	_ = rootCmd.Execute()
}
