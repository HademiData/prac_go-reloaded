package helperFunctions

func UpperCasing(input string) string {
	runes := []rune(input)

	for i := range runes {
		if isLower(runes[i]) {
			runes[i] = runes[i] - ('a' - 'A')
		}
	}
	return string(runes)
}

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}
