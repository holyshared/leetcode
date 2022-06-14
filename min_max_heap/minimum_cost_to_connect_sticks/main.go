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

func connectSticks(sticks []int) int {
	currentSticks := make(PriorityQueue, len(sticks))

	for i, length := range sticks {
		currentSticks[i] = &Item{value: length}
	}

	heap.Init(&currentSticks)

	cost := 0
	for currentSticks.Len() > 1 {
		x := heap.Pop(&currentSticks).(*Item)
		y := heap.Pop(&currentSticks).(*Item)
		combineLength := x.value + y.value
		cost += combineLength
		heap.Push(&currentSticks, &Item{value: combineLength})
	}

	return cost
}
