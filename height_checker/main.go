package main

import "sort"

func heightChecker(heights []int) int {
	cloned := make([]int, len(heights))
	copy(cloned, heights)

	sort.Sort(sort.IntSlice(cloned))

	changedCount := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != cloned[i] {
			changedCount++
		}
	}
	return changedCount
}
