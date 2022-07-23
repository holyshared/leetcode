package main


func query(index int, tree []int) int {
	result := 0
	for index >= 1 {
		result += tree[index]
		index -= index & -index
	}
	return result
}

func update(index int, value int, tree []int, size int) {
	index++
	for index < size {
		tree[index] += value
		index += index & -index
	}
}

func countSmaller(nums []int) []int {
	offset := 10000
	size := 2*10000 + 2
	tree := make([]int, size)
	result := []int{}

	for i := len(nums) - 1; i >= 0; i-- {
		smaller_count := query(nums[i]+offset, tree)
		result = append(result, smaller_count)
		update(nums[i]+offset, 1, tree, size)
	}

	for i := 0; i < len(result) / 2; i++ {
		result[i], result[len(result) - i - 1] = result[len(result) - i - 1], result[i]
	}

	return result
}
