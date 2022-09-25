package main

import "math"

func maxValues(nums []int) int {
	max := math.MinInt32
	for _, val := range nums {
		if max < val {
			max = val
		}
	}
	return max
}

func longestSubarray(nums []int) int {
	max := maxValues(nums)
	count := 0
	ans := 0
	for _, num := range nums {
		if num == max {
			count++
		} else {
			if ans < count {
				ans = count
			}
			count = 0
		}

	}
	if ans < count {
		ans = count
	}
	return ans
}
