package main

import (
	"prac_go-reloaded/utils/processors/helperFunctions"
	"testing"
)

func TestUpperCase(t *testing.T) {

	test := []struct {
		input, result string
	}{
		{
			"adewale",
			"ADEWALE"},
		{
			"paulina",
			"PAULINA"},
		{
			"",
			""},
	}

	for _, tc := range test {
		output := helperFunctions.UpperCasing(tc.input)
		if output != tc.result {
			t.Errorf("UpperCasing(%s) failed : expected %s got %s", tc.input, tc.result, output)
		}
	}

}
