package main

import (
	helperfunctions "prac_go-reloaded/utils/processors/helperFunctions"
	"testing"
)

func TestCapitalize(t *testing.T) {
	test := []struct {
		input, result string
	}{
		{
			"adewale",
			"Adewale"},
		{
			"paul",
			"Paul"},
		{
			"learn2earn",
			"Learn2earn"},
	}

	for _, tc := range test {
		output := helperfunctions.Capitalize(tc.input)
		if output != tc.result {
			t.Errorf("capitalize(%s) failed : expected %s got %s", tc.input, tc.result, output)
		}
	}
}
