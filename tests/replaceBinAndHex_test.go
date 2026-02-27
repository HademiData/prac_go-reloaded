package main

import "testing"

// Function must start with "Test" and take *testing.T
func TestReplaceHexAndBin(t *testing.T) {

	// Define a "table" of test cases
	tests := []struct {
		sample, result string
	}{
		{
			"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66  and 2  and you will see the result is 68."},
		{
			"I am 90 (hex) cars short of becoming 120 (bin) billion rich",
			"I am 144 cars short of becoming 8 billion rich"},
		{
			"she is 370 (bin) years old",
			"she is 26 years old"},
	}

	
	t.Errorf("Expected , but got ")
}
