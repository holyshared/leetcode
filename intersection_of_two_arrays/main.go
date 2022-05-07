package main

func intersection(nums1 []int, nums2 []int) []int {
	i := 0
	nums := map[int]int{}
	markedMums := map[int]bool{}
	results := []int{}

	for i <= len(nums1)-1 {
		nums[nums1[i]] = nums1[i]
		i++
	}
	i = 0
	for i <= len(nums2)-1 {
		v, ok := nums[nums2[i]]
		marked, _ := markedMums[nums2[i]]

		if ok && marked == false {
			results = append(results, v)
			markedMums[v] = true
		}
		i++
	}

	return results
}
