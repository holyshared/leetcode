package main

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for combSum := 1; combSum < target+1; combSum++ {
		for _, num := range nums {
			if combSum-num >= 0 {
				dp[combSum] += dp[combSum-num]
			}
		}
	}
	return dp[target]
}
