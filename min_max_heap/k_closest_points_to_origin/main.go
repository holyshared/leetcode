package main

import (
	"container/heap"
	"math"
)

type Item struct {
	value float64
	index int
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

func kClosest(points [][]int, k int) [][]int {
	distances := make(PriorityQueue, len(points))

	for i, point := range points {
		x, y := point[0], point[1]
		xx := math.Pow(float64(0-x), 2)
		yy := math.Pow(float64(0-y), 2)

		distance := math.Sqrt(float64(xx + yy))
		distances[i] = &Item{
			value: distance,
			index: i,
		}
	}

	heap.Init(&distances)

	results := make([][]int, k)

	for i := 0; i < k; i++ {
		item := heap.Pop(&distances).(*Item)
		results[i] = points[item.index]
	}

	return results
}
