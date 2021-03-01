package main

import (
		"testing"
)

func Test123000(t *testing.T) {
	nums1 := []int{1,2,3,0,0,0}
	m := 3
	nums2 := []int{2,5,6}
	n := 3

	// [1,2,2,3,5,6]
	merge(nums1, m, nums2, n)

	if nums1[0] != 1 {
		t.Fatalf("failed nums1 = %d", nums1)
	} 
	if nums1[1] != 2 {
		t.Fatalf("failed nums1 = %d", nums1)
	}
	if nums1[2] != 2 {
		t.Fatalf("failed nums1 = %d", nums1)
	}
	if nums1[3] != 3 {
		t.Fatalf("failed nums1 = %d", nums1)
	}
	if nums1[4] != 5 {
		t.Fatalf("failed nums1 = %d", nums1)
	}
	if nums1[5] != 6 {
		t.Fatalf("failed nums1 = %d", nums1)
	}
}

func Test1(t *testing.T) {
	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0

	// [1]
	merge(nums1, m, nums2, n)

	if nums1[0] != 1 {
		t.Fatalf("failed test %d", nums1)
	}
}

func Test0(t *testing.T) {
	nums1 := []int{0}
	m := 0
	nums2 := []int{1}
	n := 1

	// [1]
	merge(nums1, m, nums2, n)

	if nums1[0] != 1 {
		t.Fatalf("failed test %d", nums1)
	}
}
