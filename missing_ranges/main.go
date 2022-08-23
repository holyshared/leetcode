package main

import (
	"strconv"
)

// [0,1,3,50,75], lower = 0, upper = 99
// ["2","4->49","51->74","76->99"]

func missingRangeOf(nums []int) []string {
	results := []string{}

	for i := 1; i < len(nums); i++ {
		curr := nums[i - 1]
		next := nums[i]
		if next > curr + 1 {
			ns := curr + 1
			ne := next - 1
			results = append(results, missingRanges(ns, ne))
		}
	}
	return results
}

func missingRanges(start int, end int) string {
	if start == end {
		return strconv.Itoa(start)
	} else {
		return strconv.Itoa(start)+"->"+strconv.Itoa(end)
	}
}

func findMissingRanges(nums []int, lower int, upper int) []string {
	results := []string{}

	clonedNums := make([]int, len(nums))
	copy(clonedNums, nums)

	clonedNums = append(clonedNums, upper + 1)
	clonedNums = append([]int{lower-1}, clonedNums...)

	results = append(results, missingRangeOf(clonedNums)...)

	return results
}
