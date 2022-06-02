package main

import (
	"container/heap"
)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func topKFrequent(nums []int, k int) []int {
	count := map[int]int{}

	for _, v := range nums {
		current, ok := count[v]
		if ok {
			count[v] = current + 1
		} else {
			count[v] = 1
		}
	}

	pq := make(PriorityQueue, 0)

	for value, priority := range count {
		heap.Push(&pq, &Item{
			value:    value,
			priority: priority,
			index:    0,
		})
		if pq.Len() > k {
			heap.Pop(&pq)
		}
	}

	top := make([]int, k)

	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(&pq).(*Item)
		top[i] = item.value
	}

	return top
}
