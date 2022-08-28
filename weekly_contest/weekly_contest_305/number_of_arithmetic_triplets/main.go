package main


func arithmeticTriplets(nums []int, diff int) int {
	n := len(nums)

	answer := 0
	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			for k := 2; k < n; k++ {
				if nums[j] - nums[i] == diff && nums[k] - nums[j] == diff {
					answer++
				}
			}
		}
	}
	return answer
}