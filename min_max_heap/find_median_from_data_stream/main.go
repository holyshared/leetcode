package main

import (
	"container/heap"
	//	"fmt"
)

type Item struct {
	value int
}

type HighHeap []*Item
type LowHeap []*Item

func (hi HighHeap) Len() int { return len(hi) }

func (hi HighHeap) Less(i, j int) bool {
	return hi[i].value < hi[j].value
}

func (hi HighHeap) Swap(i, j int) {
	hi[i], hi[j] = hi[j], hi[i]
}

func (hi *HighHeap) Push(x interface{}) {
	item := x.(*Item)
	*hi = append(*hi, item)
}

func (hi *HighHeap) Pop() interface{} {
	old := *hi
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*hi = old[0 : n-1]
	return item
}

func (lw LowHeap) Len() int { return len(lw) }

func (lw LowHeap) Less(i, j int) bool {
	return lw[i].value > lw[j].value
}

func (lw LowHeap) Swap(i, j int) {
	lw[i], lw[j] = lw[j], lw[i]
}

func (lw *LowHeap) Push(x interface{}) {
	item := x.(*Item)
	*lw = append(*lw, item)
}

func (lw *LowHeap) Pop() interface{} {
	old := *lw
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*lw = old[0 : n-1]
	return item
}

type MedianFinder struct {
	hi HighHeap
	lo LowHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		hi: make(HighHeap, 0),
		lo: make(LowHeap, 0),
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.lo, &Item{value: num})
	heap.Push(&this.hi, this.lo[0])

	heap.Pop(&this.lo)

	if this.lo.Len() < this.hi.Len() {
		heap.Push(&this.lo, this.hi[0])
		heap.Pop(&this.hi)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.lo.Len() > this.hi.Len() {
		return float64(this.lo[0].value)
	} else {
		sum := float64(this.lo[0].value + this.hi[0].value)
		return sum * 0.5
	}
}
