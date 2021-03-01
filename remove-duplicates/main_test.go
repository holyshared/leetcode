package main

import "testing"

func Test012(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := []int{1, 2}

	resized := removeDuplicates(nums)
	if resized != 2 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}

func Test0011122334(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	expected := []int{0, 1, 2, 3, 4}

	resized := removeDuplicates(nums)
	if resized != 5 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}

func Test0(t *testing.T) {
	nums := []int{0}
	expected := []int{0}

	resized := removeDuplicates(nums)
	if resized != 1 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}

func TestEmpty(t *testing.T) {
	nums := []int{}
	expected := []int{}

	resized := removeDuplicates(nums)
	if resized != 0 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}

func Test12(t *testing.T) {
	nums := []int{1, 2}
	expected := []int{1, 2}

	resized := removeDuplicates(nums)
	if resized != 2 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}

func Test122(t *testing.T) {
	nums := []int{1, 2, 2}
	expected := []int{1, 2}

	resized := removeDuplicates(nums)
	if resized != 2 {
		t.Fatalf("resized = %d, nums = %d", resized, nums)
	}

	for i := 0; i < len(expected); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("i = %d, nums = %d", i, nums)
		}
	}
}
