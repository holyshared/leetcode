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

type QueueItem struct {
	index   int
	bricks  int
	ladders int
}

type Queue []*QueueItem

func furthestBuilding(heights []int, bricks int, ladders int) int {
	queue := make(Queue, 0)
	queue = append(queue, &QueueItem{
		index:   0,
		bricks:  bricks,
		ladders: ladders,
	})

	results := make(PriorityQueue, 0)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		nextIndex := item.index + 1
		if nextIndex > len(heights)-1 {
			return item.index
		}

		if heights[item.index] >= heights[nextIndex] {
			queue = append(queue, &QueueItem{
				index:   nextIndex,
				bricks:  item.bricks,
				ladders: item.ladders,
			})
			continue
		}

		remain := heights[item.index+1] - heights[item.index]

		if item.bricks < remain && item.ladders <= 0 {
			heap.Push(&results, &Item{
				value: item.index,
			})
			continue
		}

		if item.bricks >= remain {
			queue = append(queue, &QueueItem{
				index:   nextIndex,
				bricks:  bricks - remain,
				ladders: item.ladders,
			})
		}

		if item.ladders > 0 {
			queue = append(queue, &QueueItem{
				index:   nextIndex,
				bricks:  item.bricks,
				ladders: item.ladders - 1,
			})
		}
	}

	if results.Len() > 0 {
		item := heap.Pop(&results).(*Item)
		return item.value
	} else {
		return 0
	}
}
