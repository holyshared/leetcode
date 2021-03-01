package main

func merge(nums1 []int, m int, nums2 []int, n int)  {
	var i int
	clone := make([]int, m)

	for i = 0; i < m; i++ {
		clone[i] = nums1[i]
	}
	n1i := 0
	n2i := 0
	for i = 0; i < m + n; i++ {
		if n2i >= n || (n1i < m && clone[n1i] < nums2[n2i]) {
			nums1[i] = clone[n1i]
			n1i++
		} else {
			nums1[i] = nums2[n2i]
			n2i++
		}
	}
}
