package main

import "math"

var dirs = [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}

type Solution struct {
	maze        [][]int
	m           int
	n           int
	destination []int
	distance    [][]int
}

func (this *Solution) dfs(x, y int) {
	for _, dir := range dirs {
		nx := x + dir[0]
		ny := y + dir[1]
		count := 0

		// 有効な範囲で進めなくなるまで移動する
		for nx >= 0 && ny >= 0 && nx < this.m && ny < this.n && this.maze[nx][ny] == 0 {
			nx += dir[0]
			ny += dir[1]
			count++
		}

		// 移動位置からすすめなくなる場所までの距離より、最終地点隣の移動距離が大きい場合は
		// 移動距離を最小値に置き換える
		if this.distance[x][y]+count < this.distance[nx-dir[0]][ny-dir[1]] {
			this.distance[nx-dir[0]][ny-dir[1]] = this.distance[x][y] + count
			this.dfs(nx-dir[0], ny-dir[1]) // 最終地点から探索開始
		}
	}
}

func (this *Solution) resolve(start []int) int {
	sx, sy := start[0], start[1]
	this.distance[sx][sy] = 0
	this.dfs(sx, sy)

	dx, dy := this.destination[0], this.destination[1]
	ans := this.distance[dx][dy]
	if ans == math.MaxInt32 {
		return -1
	} else {
		return ans
	}
}

func CreateSolution(maze [][]int, destination []int) *Solution {
	m := len(maze)
	n := len(maze[0])

	distance := make([][]int, m)
	for i := 0; i < m; i++ {
		distance[i] = make([]int, n)
		for j := 0; j < n; j++ {
			distance[i][j] = math.MaxInt32
		}
	}

	return &Solution{
		maze:        maze,
		m:           m,
		n:           n,
		destination: destination,
		distance:    distance,
	}
}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	sol := CreateSolution(maze, destination)
	return sol.resolve(start)
}
