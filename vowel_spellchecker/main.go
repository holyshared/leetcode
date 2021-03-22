package main

import (
	"strings"
)

func isVowel(c rune) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	} else {
		return false
	}
}

func vowel(vowelWordDict map[string]string, query string) (string, bool) {
	runes := []rune(strings.ToLower(query))

	for i := 0; i < len(runes); i++ {
		if !isVowel(runes[i]) {
			continue
		}
		runes[i] = '*'
	}
	correct, ok := vowelWordDict[string(runes)]
	if ok {
		return correct, true
	}
	return "", false
}

func spellchecker(wordlist []string, queries []string) []string {
	sameWordDict := map[string]string{}
	nearWordDict := map[string]string{}
	vowelWordDict := map[string]string{}

	for i := 0; i < len(wordlist); i++ {
		sameWordDict[wordlist[i]] = wordlist[i]

		key := strings.ToLower(wordlist[i])
		_, ok := nearWordDict[key]
		if !ok {
			nearWordDict[key] = wordlist[i]
		}

		wildcards := []rune(key)
		for i := 0; i < len(wildcards); i++ {
			if isVowel(wildcards[i]) {
				wildcards[i] = '*'
			}
		}

		_, hasVowel := vowelWordDict[string(wildcards)]
		if !hasVowel {
			vowelWordDict[string(wildcards)] = wordlist[i]
		}
	}

	results := []string{}

	for i := 0; i < len(queries); i++ {
		q := queries[i]

		sameWord, hasSameWord := sameWordDict[q]
		if hasSameWord {
			results = append(results, sameWord)
			continue
		}

		lq := strings.ToLower(q)

		word, hasWord := nearWordDict[lq]
		if hasWord {
			results = append(results, word)
			continue
		}

		s, ok := vowel(vowelWordDict, q)
		if ok {
			results = append(results, s)
			continue
		}
		results = append(results, "")
	}

	return results
}
