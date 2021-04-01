package main

import (
	"sort"
)

func largestUniqueNumber(A []int) int {
	sort.Sort(sort.IntSlice(A))

	uniqueNumbers := map[int]int{}

	maxNum := -1
	for i := len(A) - 1; i >= 0; i-- {
		_, ok := uniqueNumbers[A[i]]
		if ok {
			maxNum = -1
			uniqueNumbers[A[i]] += 1
		} else {
			if maxNum > A[i] && maxNum != -1 {
				return maxNum
			}
			maxNum = A[i]
			uniqueNumbers[A[i]] += 1
		}
	}
	v, ok := uniqueNumbers[maxNum]
	if ok && v > 1 {
		return -1
	}
	return maxNum
}
