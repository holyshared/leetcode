package main

var move = [][]int{
	{0, -1},  // top
	{-1, 0},  // left
	{1, 0},   // right
	{0, 1},   // bottom
	{-1, -1}, // left top
	{1, -1},  // right top
	{-1, 1},  // left bottom
	{1, 1},   // right bottom
}

func shortestPathBinaryMatrix(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1
	visited := make([][]bool, m + 1)

	for i := 0; i < m + 1; i++ {
		visited[i] = make([]bool, n + 1)
	}

	queue := [][]int{{0, 0, 1}}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			next := queue[0]
			queue = queue[1:]
			x, y, d := next[0], next[1], next[2]

			if grid[x][y] != 0 {
				continue
			}

			if visited[x][y] {
				continue
			}

			if x == m && y == n {
				return d
			}

			visited[x][y] = true

			for _, mv := range move {
				nx, ny := x+mv[0], y+mv[1]
				if nx < 0 || ny < 0 || nx > m || ny > n {
					continue
				}
				queue = append(queue, []int{nx, ny, d + 1})
			}
		}
	}

	return -1
}
