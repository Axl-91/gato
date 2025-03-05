package cmd

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestClearCommand(t *testing.T) {
	testClearWithInvalidArgs(t)
	testClearAllValues(t)

	// TODO: test clear command with parameters.
	// testClearWithParameter(t)
}

// Check for invalid amount of args
func testClearWithInvalidArgs(t *testing.T) {
	testClearWithInvalidAmount(t)
	testClearWithInvalidMethod(t)
}

func testClearWithInvalidAmount(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetErr(&buf)
	rootCmd.SetArgs([]string{"clear", "port", "another"})

	_ = rootCmd.Execute()

	output := buf.String()
	expectedMsg := "accepts 0 arg(s), received 1"

	if !strings.Contains(output, expectedMsg) {
		t.Errorf("Expected output to contain %q but got %q", expectedMsg, output)
	}
}

func testClearWithInvalidMethod(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	rootCmd.SetArgs([]string{"clear", "invalid"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Error executing clear command: %v", err)
	}
	output := buf.String()

	if len(output) > 0 {
		t.Errorf("Expected empty response, instead got: %s", output)
	}
}

func testClearAllValues(t *testing.T) {
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

	// I'll use the clear command to restore all the values
	rootCmd.SetArgs([]string{"clear"})
	_ = rootCmd.Execute()

	// Check all values are the default ones
	checkValues(t, defaultValues)
}
