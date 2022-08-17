package main

import "strings"

import "strings"

var morse = []string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	ans := 0
	seen := map[string]bool{}

	for _, word := range words {
		results := make([]string, len(word))
		for i := 0; i < len(word); i++ {
			results[i] = morse[word[i]-'a']
		}
		key := strings.Join(results, "")
		if _, has := seen[key]; !has {
			ans++
			seen[key] = true
		}
	}

	return ans
}
