package processors

import (
	"strconv"
	"strings"

	"prac_go-reloaded/utils/processors/helperFunctions"
)

func FormatCasing(input string) string {
	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {

		// Formatting Casing without numbers
		if words[i] == "(cap)" || words[i] == "(cap)," && i > 0 {
			words[i-1] = helperFunctions.Capitalize(words[i-1])

			words[i] = ""

		} else if words[i] == "(up)" || words[i] == "(up)," && i > 0 {
			words[i-1] = helperFunctions.UpperCasing(words[i-1])

			words[i] = ""

		} else if words[i] == "(low)" || words[i] == "(low)," && i > 0 {
			words[i-1] = helperFunctions.LowerCasing(words[i-1])

			words[i] = ""
		}

		// Formatting Casing with Number
		var convertFunc func(string) string

		if words[i] == "(cap," || words[i] == "(up," || words[i] == "(low," {

			switch words[i] {

			case "(cap,":
				convertFunc = helperFunctions.Capitalize

			case "(up,":
				convertFunc = helperFunctions.UpperCasing

			case "(low,":
				convertFunc = helperFunctions.LowerCasing

			}

			// extracting number of words to Apply casing from words[i+1]
			num, err := strconv.Atoi(string(words[i+1][0]))

			if err == nil {
				// clear te array elem containing the number
				words[i+1] = ""
			}

			// logic to apply the conversion function on the number of specific words

			for v := 0; v < num && i-v-1 >= 0; v++ {
				prevIndex := i - v - 1
				if convertFunc != nil {
					words[prevIndex] = convertFunc(words[prevIndex])
				}
			}
			words[i] = ""
		}

	}

	var cleanWords []string

	for _, word := range words {
		if word == "" {
			continue
		} else {
			cleanWords = append(cleanWords, word)
		}

	}

	return strings.Join(cleanWords, " ")
}
