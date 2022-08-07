package main

import (
	"container/heap"
	"math"
)

type Edge struct {
	point1 int
	point2 int
	cost   int
}

type PriorityQueue []*Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Edge)
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

func minCostConnectPoints(points [][]int) int {
	if points == nil || len(points) == 0 {
		return 0
	}

	size := len(points)
	pq := make(PriorityQueue, 0)

	visited := make([]bool, size)
	result := 0
	count := size - 1

	coordinate1 := points[0]
	for j := 1; j < size; j++ {
		coordinate2 := points[j]
		cost := int(math.Abs(float64(coordinate1[0]-coordinate2[0]))) +
			int(math.Abs(float64(coordinate1[1]-coordinate2[1])))

		heap.Push(&pq, &Edge{
			point1: 0,
			point2: j,
			cost:   cost,
		})
	}
	visited[0] = true

	for len(pq) > 0 && count > 0 {
		edge := heap.Pop(&pq).(*Edge)

		point2 := edge.point2
		cost := edge.cost

		if !visited[point2] {
			result += cost
			visited[point2] = true

			for j := 0; j < size; j++ {
				if !visited[j] {
					distance := int(math.Abs(float64(points[point2][0]-points[j][0]))) + int(math.Abs(float64(points[point2][1]-points[j][1])))

					heap.Push(&pq, &Edge{
						point1: point2,
						point2: j,
						cost:   distance,
					})
				}
			}
			count--
		}
	}
	return result
}
