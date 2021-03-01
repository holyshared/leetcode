package main

func findDisappearedNumbers(nums []int) []int {
	needs := []int{}

	cache := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = true
	}

	for j := 1; j <= len(nums); j++ {
		_, ok := cache[j]
		if ok {
			continue
		}
		needs = append(needs, j)
	}
	return needs
}
