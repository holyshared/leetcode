package main

import "sort"

func minMoves2(nums []int) int {
	sort.Slice(nums, func (i, j int) bool {
		return nums[i] < nums[j]
	})

	mid := nums[len(nums)/2]

	sum := 0
	for i := 0 ; i < len(nums); i++ {
		if nums[i] > mid {
			sum += (nums[i] - mid)
		} else if nums[i] < mid {
			sum += (mid - nums[i])
		} 
	}
	return sum
}
