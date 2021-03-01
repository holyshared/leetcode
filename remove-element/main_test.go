package main

import "testing"

func Test3223(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	resized := removeElement(nums, 3)

	if resized != 2 {
		t.Fatalf("remove element faild: %d", resized)
	}

	expected := []int{2, 2}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}

func Test01223042(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	resized := removeElement(nums, 2)

	if resized != 5 {
		t.Fatalf("remove element faild: %d", resized)
	}

	expected := []int{0, 1, 4, 0, 3}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}

func Test2(t *testing.T) {
	nums := []int{2}
	resized := removeElement(nums, 3)

	if resized != 1 {
		t.Fatalf("remove element faild: %d %d", resized, nums)
	}

	expected := []int{2}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}

func Test1(t *testing.T) {
	nums := []int{1}
	resized := removeElement(nums, 1)

	if resized != 0 {
		t.Fatalf("remove element faild: %d %d", resized, nums)
	}

	expected := []int{2}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}

func Test33(t *testing.T) {
	nums := []int{3, 3}
	resized := removeElement(nums, 3)

	if resized != 0 {
		t.Fatalf("remove element faild: %d %d", resized, nums)
	}

	expected := []int{2}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}

func Test45(t *testing.T) {
	nums := []int{4, 5}
	resized := removeElement(nums, 4)

	if resized != 1 {
		t.Fatalf("remove element faild: %d %d", resized, nums)
	}

	expected := []int{5}

	for i := 0; i < len(expected)-1; i++ {
		if nums[i] == expected[i] {
			continue
		}
		t.Fatalf("remove element faild: %d, %d", len(nums), nums)
	}
}
