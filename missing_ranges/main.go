package main

import "strconv"

// [0,1,3,50,75], lower = 0, upper = 99
// ["2","4->49","51->74","76->99"]

func findMissingRanges(nums []int, lower int, upper int) []string {
	results := []string{}
	if len(nums) == 1 {
		return results
	}
	numbers := []int{}
	if lower != nums[0] {
		numbers = append(numbers, lower)
	}

	for _, num := range nums {
		numbers = append(numbers, num)
	}

	if upper != nums[len(nums)-1] {
		numbers = append(numbers, upper)
	}
	//[0,1,3,50,75,99]

	curr := numbers[0]
	for i := 1; i < len(numbers); i++ {
		next := numbers[i]

		if next-curr == 1 {
			curr = next
			continue
		}

		if i == len(numbers)-1 {
			ns := curr + 1
			results = append(results, strconv.Itoa(ns)+"->"+strconv.Itoa(numbers[i]))
		} else {
			ns := curr + 1
			ne := next - 1
			if ns == ne {
				results = append(results, strconv.Itoa(ns))
			} else {
				results = append(results, strconv.Itoa(ns)+"->"+strconv.Itoa(ne))
			}
		}
		curr = next
	}
	return results
}
