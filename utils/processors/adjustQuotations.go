package processors

import (
	"strings"
)

func AdjustQuotations(input string) string {

	words := strings.Fields(input)

	quote := "'"
	quotationStart := true

	for i := 0; i < len(words); i++ {

		if words[i] == quote {

			if quotationStart {
				words[i+1] = quote + words[i+1]
				words = append(words[:i], words[i+1:]...)
				quotationStart = false
			} else {
				words[i-1] = words[i-1] + quote
				words = append(words[:i], words[i+1:]...)

				quotationStart = true
			}
		}

	}
	return strings.Join(words, " ")
}
