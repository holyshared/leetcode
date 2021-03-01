package main

import "testing"

func Test17185461(t *testing.T) {
	expected := []int{18, 6, 6, 6, 1, -1}
	nums := replaceElements([]int{17, 18, 5, 4, 6, 1})

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}

func Test400(t *testing.T) {
	expected := []int{-1}
	nums := replaceElements([]int{400})

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}

func Test12(t *testing.T) {
	expected := []int{2, -1}
	nums := replaceElements([]int{1, 2})

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}

func Test70605x4(t *testing.T) {
	expected := []int{70605, 70605, 70605, 70605, -1}
	nums := replaceElements([]int{57010, 40840, 69871, 14425, 70605})

	for i := 0; i < len(nums); i++ {
		if nums[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}
