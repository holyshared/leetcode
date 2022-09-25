package main

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func goodIndices(nums []int, k int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)

	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := []int{}

	for i := k; i < n-k; i++ {
		if min(left[i-1], right[i+1]) >= k-1 {
			ans = append(ans, i)
		}
	}
	return ans
}
