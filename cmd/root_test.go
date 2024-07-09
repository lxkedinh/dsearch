package cmd_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/lxkedinh/dsearch/cmd"
)

var numArgTests = []struct {
	name string
	args []string
}{
	{"zero args", []string{}},
	{"one arg", []string{"foo"}},
	{"two args", []string{"foo", "bar"}},
}

func TestNumArguments(t *testing.T) {
	for _, test := range numArgTests {
		t.Run(test.name, func(t *testing.T) {
			err := cmd.ValidateArgs(cmd.RootCmd, test.args)

			switch len(test.args) {
			case 0, 2:
				if err == nil {
					t.Fatalf("Test for calling dsearch with 0 or more than 1 arguments failed. Expected error but none caught.\nArgs: %v\n", test.args)
				}
			case 1:
				if err != nil {
					t.Fatalf("test for calling dsearch with 1 argument failed.\nArgs: %v\nError: %v\n", test.args, err)
				}
			}
		})
	}
}

func TestInvalidStartDir(t *testing.T) {
	startDir := "/invaliddir"
	fmt.Println(os.Executable())
	err := cmd.ValidateStartDir(startDir)
	if err == nil {
		t.Fatalf("Invalid starting directory test did not catch expected error:\nDir: %s\n", startDir)
	}
}

var fuzzyThresholdTests = []struct {
	name      string
	threshold float64
}{
	{"valid", 0.7},
	{"invalid negative", -1.4},
	{"invalid positive", 1.3},
}

func TestInvalidFuzzyThreshold(t *testing.T) {
	for _, test := range fuzzyThresholdTests {
		t.Run(test.name, func(t *testing.T) {
			err := cmd.ValidateFuzzyThreshold(test.threshold)
			switch test.name {
			case "valid":
				if err != nil {
					t.Fatalf("Test with valid fuzzy threshold input failed.\nError: %v\n", err)
				}
			case "default":
				if err == nil {
					t.Fatalf("\"%s\" fuzzy threshold test failed. Did not error like expected\n", test.name)
				}
			}
		})
	}
}
