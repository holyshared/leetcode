package main

import "testing"

// n = 6, speed = [2,10,3,1,5,8], efficiency = [5,4,3,9,7,2], k = 2
func TestA(t *testing.T) {
	actual := maxPerformance(6, []int{2, 10, 3, 1, 5, 8}, []int{5, 4, 3, 9, 7, 2}, 2)
	if actual != 60 {
		t.Fatalf("actual = %v, expected = 60", actual)
	}
}
