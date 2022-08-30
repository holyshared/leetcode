package main

import (
	"math"
)

func edgeScore(edges []int) int {
	scores := make([]int, len(edges))

	for i := 0; i < len(edges); i++ {
		scores[edges[i]] += i
	}

	index := 0
	max := math.MinInt
	for i := 0; i < len(scores); i++ {
		if max < scores[i] {
			max = scores[i]
			index = i
		}
	}

	return index
}