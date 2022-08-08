package main

import "fmt"

func max(values ...int) int {
	max := 0
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

// 10, 9, 2, 5, 3, 7, 101, 18
//  1, 1  1  2  2  3   4   4
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	fmt.Println(dp)

	return max(dp...)
}
