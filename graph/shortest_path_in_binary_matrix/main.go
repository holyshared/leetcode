package main

var move = [][]int{
	{0, -1},   // top
	{-1, 0},  // left
	{1, 0},   // right
	{0, 1},   // bottom
	{-1, -1}, // left top
	{1, -1},  // right top
	{-1, 1},  // left bottom
	{1, 1},   // right bottom
}

type Solution struct {
	grid    [][]int
	visited [][]bool
	m       int
	n       int
	result  int
}

// 0, 1
// 1, 0

func (this *Solution) dfs(x, y, curr int) {
	if this.grid[x][y] != 0 {
		return
	}
	if this.visited[x][y] {
		return
	}

	if x == this.m-1 && y == this.n-1 {
		if this.result > curr + 1 {
			this.result = curr + 1
		}
		return
	}

	this.visited[x][y] = true
	curr += 1

	for _, v := range move {
		x1, y1 := v[0], v[1]
		nx := x+x1
		ny := y+y1
		if nx < 0 || ny < 0 || nx > this.m-1 || ny > this.n-1 {
			continue
		}
		this.dfs(nx, ny, curr)
	}

	this.visited[x][y] = false
}

func shortestPathBinaryMatrix(grid [][]int) int {
	visited := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	solution := &Solution{
		grid:    grid,
		visited: visited,
		m:       len(grid),
		n:       len(grid[0]),
		result:  999999,
	}
	solution.dfs(0, 0, 0)

	if solution.result == 999999 {
		return  - 1
	}

	return   solution.result
}
