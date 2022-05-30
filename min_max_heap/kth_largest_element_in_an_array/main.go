package main

import (
	"math"
	"strconv"
)

type MinHeap struct {
	minHeap  []int
	heapSize int
	realSize int
}

func CreateMinHeap(heapSize int) *MinHeap {
	return &MinHeap{
		minHeap:  make([]int, heapSize+1),
		heapSize: heapSize,
		realSize: 0,
	}
}

func (m *MinHeap) Add(element int) {
	m.realSize += 1
	if m.realSize > m.heapSize {
		m.realSize -= 1
		return
	}
	m.minHeap[m.realSize] = element

	index := m.realSize
	parent := index / 2

	for m.minHeap[index] < m.minHeap[parent] && index > 1 {
		temp := m.minHeap[index]
		m.minHeap[index] = m.minHeap[parent]
		m.minHeap[parent] = temp
		index = parent
		parent = index / 2
	}
}

func (m *MinHeap) Peek() int {
	return m.minHeap[1]
}

func (m *MinHeap) Pop() int {
	if m.realSize < 1 {
		return math.MaxInt
	}

	removeElement := m.minHeap[1]
	m.minHeap[1] = m.minHeap[m.realSize]
	m.realSize -= 1
	index := 1
	for index <= m.realSize/2 {
		left := index * 2
		right := (index * 2) + 1
		if m.minHeap[index] > m.minHeap[left] || m.minHeap[index] > m.minHeap[right] {
			if m.minHeap[left] < m.minHeap[right] {
				temp := m.minHeap[left]
				m.minHeap[left] = m.minHeap[index]
				m.minHeap[index] = temp
				index = left
			} else {
				temp := m.minHeap[right]
				m.minHeap[right] = m.minHeap[index]
				m.minHeap[index] = temp
				index = right
			}
		} else {
			break
		}
	}
	return removeElement

}

func (m *MinHeap) Size() int {
	return m.realSize
}

func (m *MinHeap) String() string {
	output := ""
	for i := 1; i <= m.realSize; i++ {
		output = output + strconv.Itoa(m.minHeap[i]) + ","
	}
	return output[:len(output)-1]
}

func findKthLargest(nums []int, k int) int {
	heap := CreateMinHeap(len(nums))

	for _, v := range nums {
		heap.Add(v)
		if heap.Size() > k {
			heap.Pop()
		}
	}
	return heap.Pop()
}
