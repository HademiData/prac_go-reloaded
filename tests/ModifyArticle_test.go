package main

import (
	"prac_go-reloaded/utils/processors"
	"testing"
)

func TestModifyArticle(t *testing.T) {
	test := []struct {
		sample, result string
	}{
		{"it was a egg rack that broke, a hose, a indomie, a orangotan, a umbrella, a amala",
			"it was an egg rack that broke, an hose, an indomie, an orangotan, an umbrella, an amala"},

		{"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you."},
	}

	for _, tc := range test {
		output := processors.ModifyArticle(tc.sample)
		if output != tc.result {
			t.Errorf("\n \tModifyArticle(%s) \n\n \tFAILED, Expected :  %s \n \t Got: %s", tc.sample, tc.result, output)
		}
	}
}
