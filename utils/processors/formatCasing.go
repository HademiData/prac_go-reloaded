package processors

import (
	"strings"

	"prac_go-reloaded/utils/processors/helperFunctions"
)

func FormatCasing(input string) string {
	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {

		// Formatting Casing without numbers
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = helperFunctions.Capitalize(words[i-1])

			words = append(words[:i], words[i+1:]...)

		} else if words[i] == "(up)" && i > 0 {
			words[i-1] = helperFunctions.UpperCasing(words[i-1])

			words = append(words[:i], words[i+1:]...)

		} else if words[i] == "(low)" && i > 0 {
			words[i-1] = helperFunctions.LowerCasing(words[i-1])

			words = append(words[:i], words[i+1:]...)
		}

		// Formatting Casing with Number

	}

	return strings.Join(words, " ")

}
