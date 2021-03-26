package main

import (
	"strings"
)

func wordSubsets(A []string, B []string) []string {
	dep := map[rune]bool{}
	bkeys := []rune{}
	for i := 0; i < len(B); i++ {
		for _, c := range []rune(B[i]) {
			_, ok := dep[c]
			if ok {
				continue
			}
			dep[c] = true
			bkeys = append(bkeys, c)
		}
	}

	results := []string{}
	for i := 0; i < len(A); i++ {
		matchAll := true
		for _, c := range bkeys {
			if strings.IndexRune(A[i], c) == -1 {
				matchAll = false
				break
			}
		}
		if matchAll {
			results = append(results, A[i])
		}
	}
	return results
}
