package main

import (
	"testing"
)

func TestB4341_W53341(t *testing.T) {
	boxes := []int{4, 3, 4, 1}
	warehouse := []int{5, 3, 3, 4, 1}
	actual := maxBoxesInWarehouse(boxes, warehouse)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}

func TestB12234_W3412(t *testing.T) {
	boxes := []int{1, 2, 2, 3, 4}
	warehouse := []int{3, 4, 1, 2}
	actual := maxBoxesInWarehouse(boxes, warehouse)

	if actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
}
