package main

import "testing"

/*
func Test12(t *testing.T) {
	nums := []int{1, 2}
	if validMountainArray(nums) != false {
		t.Fatal("return true")
	}
}

func Test355(t *testing.T) {
	nums := []int{3, 5, 5}
	if validMountainArray(nums) != false {
		t.Fatal("return true")
	}
}

func Test0321(t *testing.T) {
	nums := []int{0, 3, 2, 1}
	if validMountainArray(nums) != true {
		t.Fatal("return false")
	}
}

func Test012348910111211(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11}
	if validMountainArray(nums) != true {
		t.Fatal("return false")
	}
}

func Test132(t *testing.T) {
	nums := []int{1, 3, 2}
	if validMountainArray(nums) != true {
		t.Fatal("return false")
	}
}

func Test202(t *testing.T) {
	nums := []int{2, 0, 2}
	if validMountainArray(nums) != false {
		t.Fatal("return true")
	}
}
*/

func Test9876543210(t *testing.T) {
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if validMountainArray(nums) != false {
		t.Fatal("return true")
	}
}
