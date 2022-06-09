package main

import (
	"container/heap"
)

type Item struct {
	value int
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
	n := len(matrix) * len(matrix[0])
	pq := make(PriorityQueue, n)

	i := 0
	for _, row := range matrix {
		for _, value := range row {
			pq[i] = &Item{value: value}
			i++
		}
	}
	heap.Init(&pq)

	for i := 0; i < k-1; i++ {
		heap.Pop(&pq)
	}
	return heap.Pop(&pq).(*Item).value
}
