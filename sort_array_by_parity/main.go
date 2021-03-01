package main

import (
	"sort"
)

func sortArrayByParity(A []int) []int {
	var d sort.IntSlice
	d = A

	sort.Sort(d)
	i := 0
	left := 0
	right := len(A) - 1

	r := make([]int, len(A))

	for left <= right {
		if A[i]%2 == 0 {
			r[left] = A[i]
			left++
		} else {
			r[right] = A[i]
			right--
		}
		i++
	}

	return r
}
