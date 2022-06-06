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

type KthLargest struct {
	k  int
	pq PriorityQueue
}

func Constructor(k int, nums []int) KthLargest {
	pq := make(PriorityQueue, 0)
	for _, value := range nums {
		heap.Push(&pq, &Item{value: value})
	}
	for pq.Len() > k {
		heap.Pop(&pq)
	}
	return KthLargest{
		k,
		pq,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.pq, &Item{value: val})

	for this.pq.Len() > this.k {
		heap.Pop(&this.pq)
	}

	return this.pq[0].value
}
