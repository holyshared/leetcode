package main

import (
	"sort"
)

func advantageCount(A []int, B []int) []int {
	sortedA := make([]int, len(A))
	sortedB := make([]int, len(B))

	copy(sortedA, A)
	copy(sortedB, B)

	sort.Sort(sort.IntSlice(sortedA))
	sort.Sort(sort.IntSlice(sortedB))

	assigned := map[int][]int{}
	for _, v := range B {
		assigned[v] = []int{}
	}

	j := 0
	remaining := []int{}

	for _, v := range sortedA {
		if v > sortedB[j] {
			items, ok := assigned[sortedB[j]]
			if ok {
				assigned[sortedB[j]] = append(items, v)
				j++
			} else {
				assigned[sortedB[j]] = []int{sortedB[j]}
			}
		} else {
			remaining = append(remaining, v)
		}
	}

	ans := make([]int, len(B))
	for i := 0; i < len(B); i++ {
		items, ok := assigned[B[i]]

		if (ok && len(items) > 0) {
        	ans[i] = items[len(items) - 1]
			assigned[B[i]] = items[:(len(items) - 1)]
        } else {
			ans[i] = remaining[len(remaining) - 1]
			remaining = remaining[:(len(remaining) - 1)]
        }
	}

	return ans
}
