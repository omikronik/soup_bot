package cmd

import (
	"testing"
)

// Read file into []string
func TestReadWordFile(t *testing.T) {
	out, err := ReadWordFile("./testdata/input_test.txt")
	if err != nil {
		t.Error("input file not found")
	}
	if out[0] != "hello" {
		t.Error("output does not match expected file input")
	}
}
