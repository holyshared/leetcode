package main

import (
	"sort"
)

func match(word string, dict map[rune][]int) bool {
	curr := 0
	used := map[int]bool{}

	for _, c := range []rune(word) {
		positions, ok := dict[c]
		if !ok {
			return false
		}

		insert := sort.SearchInts(positions, curr)

		if insert > len(positions) - 1 {
			return false
		}

		if positions[insert] == curr {
		  _, ok := used[curr]

		  if ok {
		    if len(positions) - 1 < insert + 1 {
			  return false
			}
			curr = positions[insert + 1]
		  } else {
			curr = positions[insert]
		  }

		} else {
  		  curr = positions[insert]
		}

		used[curr] = true
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
