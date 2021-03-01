package main

import (
	"testing"
)

func Test10253(t *testing.T) {
	nums := []int{10, 2, 5, 3}
	if checkIfExist(nums) != true {
		t.Fatal("return false")
	}
}

func Test711411(t *testing.T) {
	nums := []int{7, 1, 14, 11}
	if checkIfExist(nums) != true {
		t.Fatal("return false")
	}
}

func Test31711(t *testing.T) {
	nums := []int{3, 1, 7, 11}
	if checkIfExist(nums) != false {
		t.Fatal("return true")
	}
}

func Test36(t *testing.T) {
	nums := []int{3, 6}
	if checkIfExist(nums) != true {
		t.Fatal("return false")
	}
}
func Test63(t *testing.T) {
	nums := []int{6, 3}
	if checkIfExist(nums) != true {
		t.Fatal("return false")
	}
}

func Test3(t *testing.T) {
	nums := []int{3}
	if checkIfExist(nums) != false {
		t.Fatal("return true")
	}
}

func Test4_711418(t *testing.T) {
	nums := []int{4, -7, 11, 4, 18}
	if checkIfExist(nums) != false {
		t.Fatal("return true")
	}
}
