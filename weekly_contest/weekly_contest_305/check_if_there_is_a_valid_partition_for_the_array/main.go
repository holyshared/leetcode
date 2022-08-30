package main

func validPartition(nums []int) bool {
	dp := []bool{true, false, nums[0] == nums[1], false}

	for  i := 2; i < len(nums); i++ {
		two := nums[i] == nums[i - 1]
		three := (two && nums[i] == nums[i - 2]) || (nums[i] - 1 == nums[i - 1] && nums[i] - 2 == nums[i - 2]) // ex, 1,1,1 or 1, 2, 3
		dp[(i + 1) % 4] = (dp[(i - 1) % 4] && two) || (dp[(i - 2) % 4] && three)
	}
	return dp[len(nums) % 4]
}