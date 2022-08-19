package main

func isPossible(nums []int) bool {
	subsequences := map[int]int{}
	frequency := map[int]int{}

	for _, num := range nums {
		_, has := frequency[num]
		if has {
			frequency[num]++
		} else {
			frequency[num] = 1
		}
	}

	for _, num := range nums {
		if frequency[num] == 0 {
			continue
		}

		if subsequences[num-1] > 0 {
			subsequences[num-1]--
			subsequences[num] = subsequences[num] + 1
		} else if frequency[num+1] > 0 && frequency[num+2] > 0 {
			subsequences[num+2]++
			frequency[num+1] = frequency[num+1] - 1
			frequency[num+2] = frequency[num+2] - 1
		} else {
			return false
		}
		frequency[num]--
	}

	return true
}
