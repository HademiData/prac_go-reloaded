package main

import (
	"prac_go-reloaded/utils/processors"
	"testing"
)

func TestAdjustQuotations(t *testing.T) {
	test := []struct {
		sample, result string
	}{
		{"I am exactly how they describe me: ' awesome '",
			"I am exactly how they describe me: 'awesome'"},

		{"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'"},

		{"",
			""},
	}

	for _, tc := range test {
		output := processors.AdjustQuotations(tc.sample)

		if output != tc.result {
			t.Errorf("\n \tAdjustQuotations(%s) \n\n \tFAILED, Expected :  %s \n \t Got: %s", tc.sample, tc.result, output)
		}
	}
}
