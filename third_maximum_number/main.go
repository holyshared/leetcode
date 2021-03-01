package main

import (
	"sort"
)

func thirdMax(nums []int) int {
	sort.Sort(sort.IntSlice(nums))

	l := len(nums) - 1
	rank := 0
	tmax := nums[l]
	for right := l; right >= 0 && rank <= 1; right-- {
		if tmax > nums[right] {
			tmax = nums[right]
			rank++
		}
	}
	if rank >= 2 {
		return tmax
	} else {
		return nums[l]
	}
}
