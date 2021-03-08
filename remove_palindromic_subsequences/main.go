package main

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}

	chars := []rune(s)
	revChars := []rune{}

	for i := len(chars) - 1; i >= 0; i-- {
		revChars = append(revChars, chars[i])
	}

	if s == string(revChars) {
		return 1
	} else {
		return 2
	}
}