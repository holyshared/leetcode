package main

import (
	"sort"
)

func match(word string, dict map[rune][]int) bool {
	curr := 0
	currPosition := map[rune]int{}

	for _, c := range []rune(word) {
		positions, ok := dict[c]
		if !ok {
			return false
		}

		next, _ := currPosition[c]

		// 使えるものがない
		if next >= len(positions) {
			return false
		}

		insert := sort.SearchInts(positions, curr)

		if insert > len(positions) - 1 {
			return false
		}

		curr = positions[insert]
		currPosition[c] += 1
	}

	return true
}

func numMatchingSubseq(s string, words []string) int {
	dict := map[rune][]int{}
	for i, c := range []rune(s) {
		dict[c] = append(dict[c], i)
	}

	sec := 0
	for _, word := range words {
		if match(word, dict) {
			sec += 1
		}
	}

	return sec
}
