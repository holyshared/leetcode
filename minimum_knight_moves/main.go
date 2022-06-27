package main

var offsets [][]int = [][]int{
	{1, 2},
	{1, -2},
	{2, 1},
	{2, -1},
	{-2, -1},
	{-2, 1},
	{-1, 2},
	{-1, -2},
}

func minKnightMoves(x int, y int) int {
	visited := make([][]bool, 607)
	for i := 0; i < 607; i++ {
		visited[i] = make([]bool, 607)
	}

	queue := [][]int{{0, 0}}
	steps := 0

	for len(queue) > 0 {
		currLevelSize := len(queue)
		for i := 0; i < currLevelSize; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr[0] == x && curr[1] == y {
				return steps
			}

			for _, offset := range offsets {
				mx, my := offset[0], offset[1]
				next := []int{curr[0] + mx, curr[1] + my}

				if !visited[next[0]+302][next[1]+302] {
					visited[next[0]+302][next[1]+302] = true
					queue = append(queue, next)
				}
			}
		}
		steps++
	}

	return steps
}
