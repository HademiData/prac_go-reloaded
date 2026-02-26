package main

import (
	"testing"

	"prac_go-reloaded/internal/processors"
)

// This must be named main, not main.go
func TestReplaceHexAndBin(t *testing.T) {
	// Defining a Table of Test Cases

	tests := []struct {
		sample, result string
	}{
		{
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},

		{

			"I am 90 (hex) chairs away from becoming 101 (bin) billion dollars richer",
			"I am 144 chairs away from becoming 5 billion dollars richer"},

		{
			"She is 110101 (bin) years old",
			"She is 53 years old"},
	}

	for _, tc := range tests {
		output := processors.ReplaceHexAndBin(tc.sample)

		if output != tc.result {
			t.Errorf("ReplaceHexAndBin(%s) failed: expected %s, got %s", tc.sample, tc.result, output)
		}
	}

}
