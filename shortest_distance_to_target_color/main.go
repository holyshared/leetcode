package main

import (
	"sort"
)

func shortestDistanceColor(colors []int, queries [][]int) []int {

	colorDistances := make([][]int, 3)

	for i := 0; i < len(colors); i++ {
		colorIndex := colors[i] - 1
		colorDistances[colorIndex] = append(colorDistances[colorIndex], i)
	}

	results := make([]int, len(queries))

	for i := 0; i < len(queries); i++ {
		index, color := queries[i][0], queries[i][1]

		if len(colorDistances[color - 1]) <= 0 {
			results[i] = -1
			continue
		}

		indexes := colorDistances[color - 1]
		insert := sort.SearchInts(indexes, index)

		if insert < 0 {
			insert = (insert + 1) * -1
		}

		// 一番先頭になる場合
		if insert == 0 {
			results[i] = indexes[insert] - index
		// 一番末尾になる場合
		} else if (insert == len(indexes)) {
			results[i] = index - indexes[insert - 1]
		} else {
			leftNearest := index - indexes[insert - 1]
			rightNearest := indexes[insert] - index

			if leftNearest > rightNearest {
				results[i] = rightNearest
			} else {
				results[i] = leftNearest
			}
		}
	}

	return results
}
