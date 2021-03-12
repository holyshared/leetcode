package main

import (
	"strings"
)

var chars = []rune{'0', '1'}

func pattern(s string, k int, substr []rune) bool {
	if k == 0 {
		return strings.Contains(s, string(substr))
	}
	for _, v := range chars {
		substr = append(substr, v)
		ok := pattern(s, k-1, substr)
		if !ok {
			return false
		}
		substr = substr[:(len(substr) - 1)]
	}
	return true
}

func hasAllCodes(s string, k int) bool {
	return pattern(s, k, []rune{})
}
