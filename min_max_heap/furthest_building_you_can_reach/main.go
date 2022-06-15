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

func furthestBuilding(heights []int, bricks int, ladders int) int {

	ladderAllocations := make(PriorityQueue, 0)

	for i := 0; i < len(heights)-1; i++ {
		climb := heights[i+1] - heights[i]
		if climb <= 0 {
			continue
		}
		heap.Push(&ladderAllocations, &Item{value: climb})
		if ladderAllocations.Len() <= ladders {
			continue
		}
		item := heap.Pop(&ladderAllocations).(*Item)
		bricks -= item.value

		if bricks < 0 {
			return i
		}
	}
	return len(heights) - 1
}
