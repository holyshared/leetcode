package main

import (
	"container/heap"
	"sort"
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

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) <= 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	pq := make(PriorityQueue, 0)

	heap.Push(&pq, &Item{
		value: intervals[0][1],
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pq[0].value {
			heap.Pop(&pq)
		}
		heap.Push(&pq, &Item{
			value: intervals[i][1],
		})
	}

	return pq.Len()
}
