package processors

import (
	"strings"

	"prac_go-reloaded/utils/processors/helperFunctions"
)

func FormatCasing(input string) string {

	words := strings.Fields(input)

	for i := 1; i < len(words); i++ {
		if words[i] == "(cap)" && i > 0 {
			words[i-1] = helperfunctions.Capitalize(words[i-1])
		}

	}

	return strings.Join(words, " ")

}
