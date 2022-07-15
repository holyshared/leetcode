package main

import "math"

func shortestDistanceColor(colors []int, queries [][]int) []int {
	distances := make([][]int, len(colors))

	for i := 0; i < len(distances); i++ {
		distances[i] = []int{math.MaxInt, math.MaxInt, math.MaxInt}
	}

	for i := 0; i < len(colors); i++ {
		for j := 0; j < len(colors); j++ {
			abs := 0
			if i > j {
				abs = i - j
			} else {
				abs = j - i
			}
			if abs < distances[i][colors[j]-1] {
				distances[i][colors[j]-1] = abs
			}
		}
	}

	results := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		distance := distances[queries[i][0]][queries[i][1]-1]

		if distance == math.MaxInt {
			results[i] = -1
		} else {
			results[i] = distance
		}
	}
	return results
}
