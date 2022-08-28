package main

import (
	"container/heap"
)

type PriorityQueue []int

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i] < h[j] }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func diagonalSort(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	diagonals := map[int]*PriorityQueue{}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			queue, exists := diagonals[row-col]
			if !exists {
				queue = &PriorityQueue{}
			}
			heap.Push(queue, mat[row][col])
			diagonals[row-col] = queue
		}
	}

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			queue := diagonals[row-col]
			val := heap.Pop(queue).(int)
			mat[row][col] = val
		}
	}

	return mat
}
