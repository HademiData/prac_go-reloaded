package processors

import (
	"strings"
)

func ModifyArticle(input string) string {

	vowelsandh := "aeiouh"
	words := strings.Fields(input)

	for i := 0; i < len(words); i++ {

		if i+1 < len(words) && (words[i] == "a" || words[i] == "A") {
			firstLetterOfNextWord := rune(words[i+1][0])
			if strings.ContainsRune(vowelsandh, firstLetterOfNextWord) {

				if words[i] == "a" {
					words[i] = "an"
				} else {
					words[i] = "An"
				}
			}
		}
	}
	return strings.Join(words, " ")
}
