package main

import (
	"math"
	"strconv"
)

func calcs(costs [][]int, current int, color int, memo map[string]int) int {
	if len(costs) - 1 <= current {
		return costs[current][color]
	}

	key := strconv.Itoa(current) + " " + strconv.Itoa(color)
	r, ok := memo[key] 
	if ok {
		return r
	}

	f1 := 0
	f2 := 0
	next := current + 1

	if color == 0 {
		f1 = calcs(costs, next, 1, memo)
		f2 = calcs(costs, next, 2, memo)
	} else if color == 1 {
		f1 = calcs(costs, next, 0, memo)
		f2 = calcs(costs, next, 2, memo)
	} else if color == 2 {
		f1 = calcs(costs, next, 0, memo)
		f2 = calcs(costs, next, 1, memo)
	}

	total := 0
	if f1 > f2 {
		total = costs[current][color] + f2
	} else {
		total = costs[current][color] + f1
	}
	memo[key] = total
	return total
}

func minCost(costs [][]int) int {
	min := math.MaxInt 
	memo := map[string]int{}
	for i := 0; i < 3; i++ {
		cost := calcs(costs, 0, i, memo)
		if min > cost {
			min = cost
		}
	}
	return min
}
