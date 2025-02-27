package cmd

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := os.Chdir(".."); err != nil {
		fmt.Println("Failed to change directory:", err)
		os.Exit(1)
	}

	code := m.Run()
	os.Exit(code)
}
