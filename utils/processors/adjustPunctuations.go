package processors

import (
	"prac_go-reloaded/utils/processors/helperFunctions"
	"strings"
)

func AdjustPunctuations(input string) string {

	punctuations := ".,!?:;"

	words := strings.Fields(input)

	for i := 0; i < len(words); i++ {

		// Adjusting the single character punctuations

		if len(words[i]) > 0 && strings.Contains(punctuations, words[i]) {
			words[i-1] = words[i-1] + words[i]
			words[i] = ""
		}
		// adjusting punctuations that are are the first char of a string
		if len(words[i]) > 0 && strings.Contains(punctuations, string(words[i][0])) {
			words[i-1] = words[i-1] + string(words[i][0])
			words[i] = words[i][1:]
		}

		// Adjusting Group of punctuations
		if helperFunctions.GroupOfPunctuations(words[i]) {
			words[i-1] = words[i-1] + words[i]
			words[i] = ""
		}

	}

	// store for the strings that are not empty
	var cleanWords []string

	// removing strings that are empty
	for _, word := range words {
		if word == "" {
			continue
		} else {
			cleanWords = append(cleanWords, word)
		}

	}

	return strings.Join(cleanWords, " ")

}
