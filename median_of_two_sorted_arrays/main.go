package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	total := m + n

	i := 0
	mi, ni := 0, 0
	merged := make([]int, total)

	for mi <= m-1 && ni <= n-1 {
		if nums1[mi] < nums2[ni] {
			merged[i] = nums1[mi]
			i++
			mi++
		} else if nums1[mi] > nums2[ni] {
			merged[i] = nums2[ni]
			i++
			ni++
		} else {
			merged[i] = nums1[mi]
			mi++
			i++
			merged[i] = nums2[ni]
			ni++
			i++
		}
	}

	for mi <= m-1 {
		merged[i] = nums1[mi]
		mi++
		i++
	}

	for ni <= n-1 {
		merged[i] = nums2[ni]
		ni++
		i++
	}

	p := total / 2
	if total%2 == 0 {
		return (float64(merged[p-1]) + float64(merged[p])) / 2
	} else {
		return float64(merged[p])
	}
}
