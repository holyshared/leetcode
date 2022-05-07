package main

/*
Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Explanation: [9,4] is also accepted.
*/

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	n := 0
	m := map[int]int{}
	for _, n = range nums1 {
		current, ok := m[n]
		if ok {
			m[n] = current + 1
		} else {
			m[n] = 1
		}
	}
	k := 0

	for _, n = range nums2 {
		current, _ := m[n]

		if current > 0 {
			nums1[k] = n
			m[n] = current - 1
			k++
		}
	}

	return nums1[0:k]
}
