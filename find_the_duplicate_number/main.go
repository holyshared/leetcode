package main

// [4,3,4,5,2,4,1]

func findDuplicate(nums []int) int {
	low, high := 1, len(nums)-1
	duplicate := -1

	for low <= high {
		cur := (low + high) / 2

		count := 0
		for _, num := range nums {
			if num <= cur {
				count += 1
			}
		}

		if count > cur {
			duplicate = cur
			high = cur - 1
		} else {
			low = cur + 1
		}
	}
	return duplicate
}
