package main

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	longestSequence, left, right, numZeroes := 0, 0, 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			numZeroes++
		}

		for numZeroes == 2 {
			if nums[left] == 0 {
				numZeroes--
			}
			left++
		}

		longestSequence = int(math.Max(float64(longestSequence), float64(right-left+1)))

		right++
	}
	return longestSequence
}
