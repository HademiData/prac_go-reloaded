package helperFunctions

func LowerCasing(input string) string {
	runes := []rune(input)

	for i := range runes {
		if isUpper(runes[i]) {
			runes[i] = runes[i] + ('a' - 'A')
		}
	}
	return string(runes)
}

func isUpper(r rune) bool {
	return (r <= 'A' && r >= 'Z')
}
