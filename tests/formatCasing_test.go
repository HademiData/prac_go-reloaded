package main

import (
	"prac_go-reloaded/utils/processors"
	"testing"
)

func TestFormatCasing(t *testing.T) {
	test := []struct {
		sample, result string
	}{
		{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom (up, 6) , it was the age of foolishness (cap, 6) ,  IT WAS THE (low, 3) winter of despair.",
			"It was the best of times, it was the worst of TIMES , IT WAS THE AGE OF WISDOM , It Was The Age Of Foolishness , it was the winter of despair."},

		{"",
			""},
	}

	for _, tc := range test {
		output := processors.FormatCasing(tc.sample)
		if output != tc.result {
			t.Errorf("\n \tFormatCasing(%s) \n\n \tFAILED, Expected :  %s \n \t Got: %s", tc.sample, tc.result, output)
		}
	}
}
