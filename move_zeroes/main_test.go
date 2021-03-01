package main

import "testing"

func Test010312(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	expected := []int{1, 3, 12, 0, 0}
	moveZeroes(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}

func Test01(t *testing.T) {
	nums := []int{0, 1}
	expected := []int{1, 0}
	moveZeroes(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}

func Test10(t *testing.T) {
	nums := []int{1, 0}
	expected := []int{1, 0}
	moveZeroes(nums)

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}
