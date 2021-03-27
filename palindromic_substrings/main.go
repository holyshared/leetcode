package main

func isPalindromic(chars []rune) bool {
	if len(chars) == 1 {
		return true
	}
	left := 0
	right := len(chars) - 1

	for left < right {
		if chars[left] != chars[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func countSubstrings(s string) int {
	chars := []rune(s)
	l := len(chars)

	count := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if isPalindromic(chars[i:j]) {
				count++
			}
		}
	}
	return count
}
