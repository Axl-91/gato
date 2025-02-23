package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

// Function to simulate the execution of a command
func executeCommand(cmd *cobra.Command, args []string) (string, error) {
	// Capture the output of the command
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	cmd.SetErr(&buf)

	// Set the arguments for the command
	cmd.SetArgs(args)

	// Execute the command
	err := cmd.Execute()

	// Return the captured output and any error
	return buf.String(), err
}

func TestRootCommand(t *testing.T) {
	output, err := executeCommand(rootCmd, []string{})
	if err != nil {
		t.Fatalf("Error executing root command: %v", err)
	}

	expected := "Gato is an HTTP request client CLI similar to cURL\n"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q but got %q", expected, output)
	}
}
