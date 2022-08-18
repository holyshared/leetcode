package main

import (
	"sort"
)

func minSetSize(arr []int) int {
	counts := map[int]int{}
	for _, v := range arr {
		count, has := counts[v]
		if !has {
			counts[v] = 1
		} else {
			counts[v] = count + 1
		}
	}

	results := []int{}
	for _, v := range counts {
		results = append(results, v)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i] > results[j]
	})

	set := 0
	size := 0
	halfSize := len(arr) / 2

	for i := 0; i < len(results); i++ {
		size += results[i]
		set++
		if size >= halfSize {
			break
		}
	}

	return set
}
