package main

func findErrorNums(nums []int) []int {
	dep := 0
	numbers := map[int]bool{}

	for i := 0; i < len(nums); i++ {
		_, ok := numbers[nums[i]]
		if !ok {
			numbers[nums[i]] = true
		} else {
			dep = nums[i]
		}
	}
	_, lostLeft := numbers[dep-1]
	if !lostLeft {
		if dep-1 > 0 {
			return []int{dep, dep - 1}
		}
	}

	_, lostRight := numbers[dep+1]
	if !lostRight {
		return []int{dep, dep + 1}
	}

	var j int
	for j = 1; j < len(numbers); j++ {
		_, find := numbers[j]
		if find {
			continue
		}
		return []int{dep, j}
	}

	return []int{dep, j + 1}
}
