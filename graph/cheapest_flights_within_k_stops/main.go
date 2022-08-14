package main

import (
	"math"
	"strconv"
)

type Solution struct {
	adjMatrix [][]int
	memo map[string]int64
}

func (this *Solution) findShortest(node, stops, dst, n int) int64 {
	if node == dst {
		return 0
	}

	if stops < 0 {
		return math.MaxInt32
	}

	key := strconv.Itoa(node) + "-" + strconv.Itoa(stops)
	val, has := this.memo[key]

	if has {
		return val
	}

	var ans int64 = math.MaxInt32
	for neighbor := 0; neighbor < n; neighbor++ {
		weight := this.adjMatrix[node][neighbor]
		if weight > 0 {
			newVal := this.findShortest(neighbor, stops - 1, dst, n) + int64(weight)
			if newVal < ans {
				ans = newVal
			}
		}  
	}
	this.memo[key] = ans
	return ans
}

func (this *Solution) findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	for _, flight := range flights {
		this.adjMatrix[flight[0]][flight[1]] = flight[2]
	}

	ans := this.findShortest(src, k, dst, n)
	if ans >= math.MaxInt32 {
		return -1
	} else {
		return int(ans)
	}
}


func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	adjMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		adjMatrix[i] = make([]int, n)
	}

	sol := &Solution{
		adjMatrix: adjMatrix,
		memo: map[string]int64{},
	}
	return sol.findCheapestPrice(n, flights, src, dst, k)
}
