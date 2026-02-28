package main

import (
	"prac_go-reloaded/utils/processors/helperFunctions"
	"testing"
)

func TestLowerCasing(t *testing.T) {
	test := []struct {
		input, result string
	}{
		{
			"loWer",
			"lower"},
		{
			"LEARN2EARN",
			"learn2earn"},
		{
			"HADEMI",
			"hademi"},
		{
			"",
			""},
	}

	for _, tc := range test {
		output := helperFunctions.LowerCasing(tc.input)
		if output != tc.result {
			t.Errorf("LowerCasing(%s) failed : expected %s got %s", tc.input, tc.result, output)
		}

	}
}
