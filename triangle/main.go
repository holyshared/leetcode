package main

import "math"

func minimumTotal(triangle [][]int) int {
	prevRow := triangle[0]

	for row := 1; row < len(triangle); row++ {
		currRow := []int{}

		for col := 0; col <= row; col++ {
			smallestAbove := math.MaxInt64
			if col > 0 {
				smallestAbove = prevRow[col-1]
			}
			if col < row {
				if smallestAbove > prevRow[col] {
					smallestAbove = prevRow[col]
				}
			}
			currRow = append(currRow, smallestAbove+triangle[row][col])
		}
		prevRow = currRow
	}

	min := math.MaxInt64
	for _, val := range prevRow {
		if min > val {
			min = val
		}
	}

	return min
}
