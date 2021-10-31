package main

func _search(nums []int, left, right int) int {
	if left == right {
		return left
	}
	mid := (left + right) / 2
	if nums[mid] > nums[mid+1] {
		return _search(nums, left, mid)
	} else {
		return _search(nums, mid+1, right)
	}
}

func findPeakElement(nums []int) int {
	return _search(nums, 0, len(nums)-1)
}
