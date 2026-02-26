package processors

import (
	//"fmt"
	"strconv"
	"strings"
)

func ReplaceHexAndBin(text string) string {

	words := strings.Fields(text)
	// fmt.Println(words)

	for i := 1; i < len(words); i++ {

		if words[i] == "(hex)" && i > 0 {

			decimalValue, err := strconv.ParseInt(words[i-1], 16, 64)

			if err == nil {
				words[i-1] = strconv.Itoa(int(decimalValue))

				words = append(words[:i], words[i+1:]...)
			}
		}
		if words[i] == "(bin)" && i > 0 {

			decimalValue, err := strconv.ParseInt(words[i-1], 2, 64)

			if err == nil {
				words[i-1] = strconv.Itoa(int(decimalValue))

				words = append(words[:i], words[i+1:]...)

			}
		}
	}
	return strings.Join(words, " ")
}
