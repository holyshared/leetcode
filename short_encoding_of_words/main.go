package main

import (
	"strings"
)

// ["time", "me", "bell"]
func minimumLengthEncoding(words []string) int {
	depWords := map[string]bool{}
	uniqueWords := []string{}

	var i int = 0
	for i = 0; i < len(words); i++ {
		_, ok := depWords[words[i]]
		if ok {
			continue
		}
		depWords[words[i]] = true
		uniqueWords = append(uniqueWords, words[i])
	}

	skipWords := map[int]bool{}

	for i := 0; i < len(uniqueWords); i++ {
		for j := 0; j < len(uniqueWords); j++ {
			if i == j {
				continue
			}
			if !strings.HasSuffix(uniqueWords[i], uniqueWords[j]) {
				continue
			}
			skipWords[j] = true
		}
	}

	refLen := 0
	for k := 0; k < len(uniqueWords); k++ {
		_, ok := skipWords[k]
		if ok {
			continue
		}
		refLen += len(uniqueWords[k]) + 1
	}

	return refLen
}
