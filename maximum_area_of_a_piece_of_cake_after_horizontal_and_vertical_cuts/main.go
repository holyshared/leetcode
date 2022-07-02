package main

import (
	"sort"
)

var mod int64 = 1000000007

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Slice(horizontalCuts, func(i, j int) bool {
		return horizontalCuts[i] < horizontalCuts[j]
	})
	sort.Slice(verticalCuts, func(i, j int) bool {
		return verticalCuts[i] < verticalCuts[j]
	})
	var maxHeight, maxWidth, diff, curr int64
	maxHeight, maxWidth = 0, 0
	curr = int64(h)
	for i := len(horizontalCuts) - 1; i >= 0; i-- {
		diff = curr - int64(horizontalCuts[i])
		curr = int64(horizontalCuts[i])
		if diff > maxHeight {
			maxHeight = diff
		}
	}
	if curr > maxHeight {
		maxHeight = curr
	}

	curr = int64(w)
	for i := len(verticalCuts) - 1; i >= 0; i-- {
		diff := curr - int64(verticalCuts[i])
		curr = int64(verticalCuts[i])
		if diff > maxWidth {
			maxWidth = diff
		}
	}
	if curr > maxWidth {
		maxWidth = curr
	}

	return int((maxHeight * maxWidth) % mod)
}
