package cmd

import (
	"testing"
)

func TestBaseCommand(t *testing.T) {
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Base command failed: %v", err)
	}
}
