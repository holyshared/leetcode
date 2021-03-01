package main

func removeDuplicates(nums []int) int {
	l := len(nums)

	if l <= 0 {
		return l
	}

	ip := 0

	for i := 1; i < l; i++ {
		if nums[i] != nums[ip] {
			ip++
			nums[ip] = nums[i]
		}
	}

	return ip + 1
}
