package main

import (
	"container/heap"
	"sort"
)

type Item struct {
	speed      int
	efficiency int
}

type PriorityQueue []int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(int)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	candidates := make([]*Item, n)
	for i := 0; i < n; i++ {
		candidates[i] = &Item{speed: speed[i], efficiency: efficiency[i]}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].efficiency > candidates[j].efficiency
	})

	speedSum, perf := 0, 0
	speedHeap := PriorityQueue{}

	for _, p := range candidates {
		currEfficiency := p.efficiency
		currSpeed := p.speed

		if speedHeap.Len() > k-1 {
			speed := heap.Pop(&speedHeap).(int)
			speedSum -= speed
		}
		heap.Push(&speedHeap, currSpeed)

		speedSum += currSpeed
		perf = max(perf, speedSum*currEfficiency)
	}
	return perf % 1000000007
}
