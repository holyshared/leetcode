package main

import (
	"container/heap"
	"math"
)

var dirs = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

type Item struct {
	x          int
	y          int
	difference int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].difference < pq[j].difference
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func isValid(x, y, m, n int) bool {
	return x >= 0 && y >= 0 && x < m && y < n
}

func diffAbs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])

	visited := make([][]bool, m)
	differences := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		differences[i] = make([]int, n)
		for j := 0; j < n; j++ {
			differences[i][j] = math.MaxInt32
			visited[i][j] = false
		}
	}
	differences[0][0] = 0

	queue := make(PriorityQueue, 0)
	heap.Push(&queue, &Item{x: 0, y: 0, difference: differences[0][0]})

	for queue.Len() > 0 {
		curr := heap.Pop(&queue).(*Item)
		visited[curr.x][curr.y] = true

		if curr.x == m-1 && curr.y == n-1 {
			return curr.difference
		}

		for _, dir := range dirs {
			nx, ny := curr.x+dir[0], curr.y+dir[1]

			if isValid(nx, ny, m, n) && !visited[nx][ny] {
				currentDifference := diffAbs(heights[nx][ny], heights[curr.x][curr.y])
				maxDifference := max(currentDifference, differences[curr.x][curr.y])

				if differences[nx][ny] > maxDifference {
					differences[nx][ny] = maxDifference
					heap.Push(&queue, &Item{x: nx, y: ny, difference: maxDifference})
				}
			}
		}
	}

	return differences[m-1][n-1]
}
