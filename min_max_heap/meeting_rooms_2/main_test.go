package main

import (
	"testing"
)

func TestNumsK2(t *testing.T) {
	actual := minMeetingRooms([][]int{
		[]int{0, 30},
		[]int{5, 10},
		[]int{15, 20},
	})
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestNumsK1(t *testing.T) {
	actual := minMeetingRooms([][]int{
		[]int{7, 10},
		[]int{2, 4},
	})
	if actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
}

func TestNums3(t *testing.T) {
	actual := minMeetingRooms([][]int{
		[]int{1, 5},
		[]int{8, 9},
		[]int{8, 9},
	})
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}

}
