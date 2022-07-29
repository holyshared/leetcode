package main

var dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
	freshOrange := 0
	queue := [][]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				freshOrange++
			} else if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	minute := 0

	for len(queue) > 0 && freshOrange > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			x, y := curr[0], curr[1]
			queue = queue[1:]

			if grid[x][y] == 2 {
				for _, v := range dirs {
					nx, ny := x+v[0], y+v[1]
					if nx < 0 || ny < 0 || nx > len(grid) - 1 || ny > len(grid[0]) - 1 {
						continue
					}
					if grid[nx][ny] == 1 {
						grid[nx][ny] = 2
						freshOrange--
						if freshOrange <= 0 {
							break
						}
						queue = append(queue, []int{x+v[0], y+v[1]})
					}
				}
			}
		}
		minute++
	}

	if freshOrange <= 0 {
		return minute
	} else {
		return -1
	}
}
