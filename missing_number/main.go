package main

// [3,0,1]
func missingNumber(nums []int) int {
	n := len(nums)
	numbers := map[int]bool{}

	for i := 0; i < n; i++ {
		numbers[nums[i]] = true
	}

	j := 0
	for j = 0; j < n; j++ {
		_, next := numbers[j]
		if !next {
			return j
		}
	}

	return j
}
