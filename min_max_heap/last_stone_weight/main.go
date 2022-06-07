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
	return pq[i].value > pq[j].value
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


func lastStoneWeight(stones []int) int {
	pq := make(PriorityQueue, len(stones))

	for i, weight := range stones {
		pq[i] = &Item{ value: weight }
	}

	heap.Init(&pq)

	for pq.Len() > 1 {
		first := heap.Pop(&pq).(*Item)
		second := heap.Pop(&pq).(*Item)

		if first.value < second.value {
			heap.Push(&pq, &Item{ value: second.value - first.value })
		} else {
			heap.Push(&pq, &Item{ value: first.value - second.value })
		}
	}
	last := heap.Pop(&pq).(*Item)
	return last.value
}