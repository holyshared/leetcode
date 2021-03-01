package main

import "testing"

func Test3124(t *testing.T) {
	nums := []int{3, 1, 2, 4}
	expected := []int{2, 4, 1, 3}
	actual := sortArrayByParity(nums)

	for i := 0; i < len(nums); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("faild %d", nums)
		}
	}
}
