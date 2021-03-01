package main

func moveZeroes(nums []int) {
	i := 0
	positions := []int{}

	for i = 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		positions = append(positions, i)
	}

	pi := 0
	l := len(positions)
	for i = 0; i < len(nums); i++ {
		if i < l {
			nums[i] = nums[positions[pi]]
			pi++
		} else {
			nums[i] = 0
		}
	}
}
