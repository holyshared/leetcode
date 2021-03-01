package main

func removeElement(nums []int, val int) int {
	l := len(nums) - 1
	left, right, rm := 0, l, 0
	if len(nums) <= 0 {
		return 0
	}

	for left <= right {
		if nums[left] == val {
			nums[left], nums[right] = nums[right], nums[left]
			right--
			rm++
		} else {
			left++
		}
	}
	return len(nums) - rm
}
