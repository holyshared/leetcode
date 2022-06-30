package main

func runningSum(nums []int) []int {
	results := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		results = append(results, results[i - 1] + nums[i])
	}
	return results
}
