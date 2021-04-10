package main

import (
	"math"
)

var dirs [][]int = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

type Solution struct {
	matrix [][]int
	m      int
	n      int
}

func NewSolution(matrix [][]int) Solution {
	m, n := len(matrix), len(matrix[0])
	return Solution{
		matrix: matrix,
		m:      m,
		n:      n,
	}
}

func (this *Solution) dfs(i int, j int, cache [][]int) int {
	if cache[i][j] != 0 {
		return cache[i][j]
	}

	for _, d := range dirs {
		x, y := i+d[0], j+d[1]
		positionInMatrix := 0 <= x && x < this.m && 0 <= y && y < this.n
		if !positionInMatrix {
			continue
		}
		nextGreaterThanCurrent := this.matrix[x][y] > this.matrix[i][j]
		if !nextGreaterThanCurrent {
			continue
		}
		cache[i][j] = int(math.Max(float64(cache[i][j]), float64(this.dfs(x, y, cache))))
	}
	cache[i][j]++
	return cache[i][j]
}

func (this *Solution) Calculate() int {
	if len(this.matrix) == 0 {
		return 0
	}

	cache := make([][]int, this.m)
	for i := 0; this.m > i; i++ {
		cache[i] = make([]int, this.n)
	}

	ans := 0
	for i := 0; i < this.m; i++ {
		for j := 0; j < this.n; j++ {
			ans = int(math.Max(float64(ans), float64(this.dfs(i, j, cache))))
		}
	}
	return ans
}

func longestIncreasingPath(matrix [][]int) int {
	solution := NewSolution(matrix)
	return solution.Calculate()
}
