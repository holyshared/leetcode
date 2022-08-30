package main

var dirs = [][]int{
	{-1, 0}, // left
	{1, 0}, // rigth
	{0, -1}, // yop
	{0, 1}, //bottom

	{-1, -1}, // lt
	{1, -1}, // rt
	{-1, 1}, // lb
	{1, 1}, // rb
	} 


func maxAt(grid [][]int, i, j int) int {
	max := grid[i][j]
	m, n := len(grid), len(grid[0])
 	for _, dir := range dirs {
		x, y := dir[0], dir[1]
		nx, ny :=  x + i, y + j
		if nx < 0 || nx > m || ny < 0 || ny > n {
			continue
		}

		if max < grid[nx][ny] {
			max = grid[nx][ny]
		}
	}
	return max
}

func largestLocal(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
 
	maxLocal := make([][]int, m- 2)

	for i := 0; i < m- 2; i++ {
		maxLocal[i] = make([]int, n-2)
	}

	for i := 0;  i < m- 2; i++ {
		for j := 0; j < n- 2; j++ {
			maxLocal[i][j] = maxAt(grid, i + 1, j + 1)
		}
	}
	return maxLocal

}