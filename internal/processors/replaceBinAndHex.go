package processors

import (
	//"fmt"
	"strconv"
	"strings"
)

func ReplaceHexAndBin(text string) string {

	words := strings.Fields(text)

	for i := 1; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {

			decimalValue, err := strconv.ParseInt(words[i-1], 16, 64)

			if err == nil {
				words[i-1] = strconv.FormatInt(decimalValue, 10)

				words[i] = ""
			}
		} else if words[i] == "(bin)" && i > 0 {

			decimalValue, err := strconv.ParseInt(words[i-1], 2, 64)

			if err == nil {
				words[i-1] = strconv.FormatInt(decimalValue, 10)

				words[i] = ""

			}
		}
	}
	return strings.Join(words, " ")
}
