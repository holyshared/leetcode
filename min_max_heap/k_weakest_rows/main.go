package main

import (
	"container/heap"
)

type Item struct {
	value int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].value == pq[j].value {
		return pq[i].index < pq[j].index
	}
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

func kWeakestRows(mat [][]int, k int) []int {
	pq := make(PriorityQueue, len(mat))

	for i, row := range mat {
		solider := 0
		for _, sorc := range row {
			if sorc != 1 {
				continue
			}
			solider += 1
		}
		pq[i] = &Item{
			value: solider,
			index: i,
		}
	}
	heap.Init(&pq)

	results := make([]int, k)

	for i := 0; i < k; i++ {
		results[i] = heap.Pop(&pq).(*Item).index
	}

	return results
}
