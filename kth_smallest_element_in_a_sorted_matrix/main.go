package main

import (
	"container/heap"
)

type Item struct {
	value    int
	row      int
	col      int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
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
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}



func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	size := n
	if n > k {
		size = k
	}

	pq := make(PriorityQueue, size)

	for  row := 0; row < size; row++ {
		pq[row] = &Item{value: matrix[row][0], row: row, col: 0, }
	}
	heap.Init(&pq)

	element := pq[0]
	for k > 0 {
		element = heap.Pop(&pq).(*Item)
		row, col := element.row, element.col

		if col < n - 1 {
			heap.Push(&pq, &Item{value: matrix[row][col + 1], row: row, col: col + 1})
		}
		k--
	}
	return element.value
}