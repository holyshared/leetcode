package main

import "math"

func minimumSubarraysRequired(nums []int, maxSumAllowed int) int {
	currentSum, splitsRequired := 0, 0

	for _, element := range nums {
		if currentSum+element <= maxSumAllowed {
			currentSum += element
		} else {
			currentSum = element
			splitsRequired++
		}
	}

	return splitsRequired + 1
}

func splitArray(nums []int, m int) int {
	sum, maxElement := 0, math.MinInt32
	for _, element := range nums {
		sum += element
		if maxElement < element {
			maxElement = element
		}
	}

	left, right, minimumLargestSplitSum := maxElement, sum, 0
	for left <= right {
		maxSumAllowed := left + (right-left)/2

		if minimumSubarraysRequired(nums, maxSumAllowed) <= m {
			right = maxSumAllowed - 1
			minimumLargestSplitSum = maxSumAllowed
		} else {
			left = maxSumAllowed + 1
		}
	}

	return minimumLargestSplitSum
}
