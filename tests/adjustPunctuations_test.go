package main

import (
	"prac_go-reloaded/utils/processors"
	"testing"
)

func TestAdjustPunctuations(t *testing.T) {
	test := []struct {
		sample, result string
	}{
		{"I was sitting over there ,and ... then BAMM !?",
			"I was sitting over there, and... then BAMM!?"},

		{"I was sitting over there ,and ... then BAMM !...",
			"I was sitting over there, and... then BAMM!..."},

		{"I am exactly how they describe me :Mean",
			"I am exactly how they describe me: Mean"},

		{"I am exactly how they describe me ,mean",
			"I am exactly how they describe me, mean"},
	}

	for _, tc := range test {
		output := processors.AdjustPunctuations(tc.sample)

		if output != tc.result {
			t.Errorf("\n \tAdjustPunctuations(%s) \n\n \tFAILED, Expected :  %s \n \t Got: %s", tc.sample, tc.result, output)
		}
	}

}
