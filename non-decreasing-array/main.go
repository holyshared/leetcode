package main

func checkPossibility(nums []int) bool {
	numViolations := 0
	for i := 1; i < len(nums); i++ {

		if nums[i-1] > nums[i] {

			if numViolations == 1 {
				return false
			}

			numViolations++

			if i < 2 || nums[i-2] <= nums[i] {
				nums[i-1] = nums[i]
			} else {
				nums[i] = nums[i-1]
			}
		}
	}

	return true
}
