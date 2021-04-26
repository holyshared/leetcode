package main

import "sort"

func isReachable(buildingIndex int, heights []int, bricks int, ladders int) bool {
	climbs := []int{}

	for i := 0; i < buildingIndex; i++ {
		h1 := heights[i]
		h2 := heights[i+1]
		if h2 <= h1 {
			continue
		}
		climbs = append(climbs, h2-h1)
	}
	sort.Sort(sort.IntSlice(climbs))

	for _, climb := range climbs {
		if climb <= bricks {
			bricks -= climb
		} else if ladders >= 1 {
			ladders -= 1
		} else {
			return false
		}
	}
	return true
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	lo := 0
	hi := len(heights) - 1
	for lo < hi {
		mid := lo + (hi-lo+1)/2
		if isReachable(mid, heights, bricks, ladders) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return hi
}
