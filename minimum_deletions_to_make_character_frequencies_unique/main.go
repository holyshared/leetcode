package main

import (
	"sort"
)

func minDeletions(s string) int {
	chars := []rune(s)
	frequency := make([]int, 26)

	for i := 0; i < len(s); i++ {
		frequency[chars[i]-'a'] += 1
	}

	sort.Slice(frequency, func(i, j int) bool {
		return frequency[i] < frequency[j]
	})

	deleteCount := 0
	maxFreqAllowed := len(s)

	for i := 25; i >= 0 && frequency[i] > 0; i-- {
		if frequency[i] > maxFreqAllowed {
			deleteCount += frequency[i] - maxFreqAllowed
			frequency[i] = maxFreqAllowed
		}

		if frequency[i]-1 > 0 {
			maxFreqAllowed = frequency[i] - 1
		} else {
			maxFreqAllowed = 0
		}
	}

	return deleteCount
}
