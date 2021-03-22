package main

import (
	"math"
)

func smallestCommonElement(mat [][]int) int {
	rows := len(mat)
	lowColumns := math.MaxInt64

	for i := 0; i < rows; i++ {
		c := len(mat[i])
		if lowColumns > c {
			lowColumns = c
		}
	}

	numCounts := map[int]int{}

	for j := 0; j < lowColumns; j++ {
		for i := 0; i < rows; i++ {
			tv := mat[i][j]
			v, ok := numCounts[tv]
			if ok {
				v++
				if v >= rows {
					return tv
				}
				numCounts[tv] = v
			} else {
				numCounts[tv]++
			}
		}

	}

	return -1
}
