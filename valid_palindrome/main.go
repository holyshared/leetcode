package main

import (
	"strings"
)

func isLetter(c rune) bool {
	val := c - 'a'
	if val < 0 || val > 25 {
		return false
	}
	return true
}

func isDigit(c rune) bool {
	return c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9'
}

func isPalindrome(s string) bool {
	chars := []rune{}

	for _, c := range strings.ToLower(s) {
		if !isLetter(c) && !isDigit(c) {
			continue
		}
		chars = append(chars, c)
	}

	if len(chars) <= 0 {
		return true
	}

	l := len(chars)
	a := make([]rune, l)
	for i := 0; i <= l/2; i++ {
		a[i] = chars[l - i - 1]
		a[l - i - 1] = chars[i]
	}

	return string(a) == string(chars)
}