package cmd

import (
	"fmt"
	"os"
	"testing"
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
			err := ValidateArgs(rootCmd, test.args)

			switch len(test.args) {
			case 0, 2:
				if err == nil {
					t.Fatalf("Test for calling dsearch with 0 or more than 1 arguments failed.\nArgs: %v", test.args)
				}
			case 1:
				if err != nil {
					t.Fatalf("test for calling dsearch with 1 argument failed.\nArgs: %v\nError: %v", test.args, err)
				}
			}
		})
	}
}

func TestSearchPath(t *testing.T) {
	path := "foobar"
	err := ValidateSearchPath(path)
	if err != nil {
		t.Fatalf("Validating search path test failed\nPath: \"%s\"\nError: %v", path, err)
	}
}

func TestInvalidStartDir(t *testing.T) {
	startDir := "/invaliddir"
	fmt.Println(os.Executable())
	err := ValidateStartDir(startDir)
	if err == nil {
		t.Fatalf("Invalid starting directory test did not error:\nDir: %s", startDir)
	}
}
