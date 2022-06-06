package main

import (
	"testing"
)

func TestNums4582K3(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	actual := kthLargest.Add(3)
	if actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}

	actual = kthLargest.Add(5)
	if actual != 5 {
		t.Fatalf("actual = %d, expected = 5", actual)
	}

	actual = kthLargest.Add(10)
	if actual != 5 {
		t.Fatalf("actual = %d, expected = 5", actual)
	}

	actual = kthLargest.Add(9)
	if actual != 8 {
		t.Fatalf("actual = %d, expected = 8", actual)
	}

	actual = kthLargest.Add(4)
	if actual != 8 {
		t.Fatalf("actual = %d, expected = 8", actual)
	}
}
