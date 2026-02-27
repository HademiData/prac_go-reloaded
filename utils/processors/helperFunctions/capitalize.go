package helperfunctions

func Capitalize(word string) string {

	runes := []rune(word)

	if runes[0] >= 'a' && runes[0] <= 'z' {
		runes[0] = runes[0] - ('a'-'A')
	}
	return string(runes)
}
