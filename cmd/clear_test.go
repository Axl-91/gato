package cmd

import (
	"strconv"
	"testing"
)

func TestClearCommand(t *testing.T) {
	defaultValues := []string{"http://127.0.0.1", "", "8000", "GET", ""}
	newValues := []string{"http://google.com", "/api/storage", "4000", "POST", "body.json"}

	newPort, _ := strconv.Atoi(newValues[2])

	// Check that the values are the default ones.
	checkValues(t, defaultValues)

	// Change the values and use check command to validate they were changed
	setValues(newValues[0], newValues[1], newPort, newValues[3], newValues[4])
	checkValues(t, newValues)

	// I'll use the clear command to restore all the values
	rootCmd.SetArgs([]string{"clear"})
	_ = rootCmd.Execute()

	// Check all values are the default ones
	checkValues(t, defaultValues)

	// TODO: test clear command with every parameter.
	// The clear should only change that value.
}
