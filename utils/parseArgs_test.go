package utils

import (
	"os"
	"testing"
)

func TestParseArgs(t *testing.T) {
	originalArgs := os.Args

	os.Args = []string{"command"}
	input, err := ParseArgs()
	if input != "" || err != "Please provide an input file" {
		t.Errorf("Expected input: '', error: 'Please provide an input file', got input: '%s', error: '%s'", input, err)
	}
	os.Args = []string{"command", "input.txt"}
	input, err = ParseArgs()
	if input != "input.txt" || err != "" {
		t.Errorf("Expected input: 'input.txt', error: '', got input: '%s', error: '%s'", input, err)
	}
	os.Args = []string{"command", "input.txt", "extra_arg"}
	input, err = ParseArgs()
	if input != "" || err != "too many arguments" {
		t.Errorf("Expected input: '', error: 'too many arguments', got input: '%s', error: '%s'", input, err)
	}
	os.Args = originalArgs
}
