package main

import "math"

func abs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	totDist, d := 0, math.MinInt64
	for _, nut := range nuts {
		totDist += (distance(nut, tree) * 2)
		deductiableMax := distance(nut, tree) - distance(nut, squirrel)
		if deductiableMax > d {
			d = deductiableMax
		}
	}
	return totDist - d
}
