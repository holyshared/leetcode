package main

func runningSum(nums []int) []int {
	curr := 0
	results := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		results[i] = curr
	}

	return results
}
