package helperFunctions

import (
	"strings"
)

func GroupOfPunctuations(input string) bool {
	punctuations := ".,!?:;"

	for _, char := range input {
		if !strings.ContainsRune(punctuations, char) {
			return false
		}
	}
	return true
}
