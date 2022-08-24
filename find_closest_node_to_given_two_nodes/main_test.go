package main

import "testing" 

func TestA(t *testing.T) {
	actual := closestMeetingNode([]int{2,2,3,-1}, 0, 1)
	if actual != 2 {
		t.Fatalf("actual = %v, expected = 2", actual)
	}
}

func TestB(t *testing.T) {
	actual := closestMeetingNode([]int{1,2,-1}, 0, 2)
	if actual != 2 {
		t.Fatalf("actual = %v, expected = 2", actual)
	}
}

// 0 -> 1 <- 2 
func TestC(t *testing.T) {
	actual := closestMeetingNode([]int{1,-1,1}, 1, 2)
	if actual != 1 {
		t.Fatalf("actual = %v, expected = 1", actual)
	}
}

// 0 -> 1 <- 2 
func TestD(t *testing.T) {
	actual := closestMeetingNode([]int{1,-1,1}, 0, 1)
	if actual != 1 {
		t.Fatalf("actual = %v, expected = 1", actual)
	}
}

// 0 -> 1  2 <- 3
func TestE(t *testing.T) {
	actual := closestMeetingNode([]int{1,-1,-1,2}, 0, 3)
	if actual != -1 {
		t.Fatalf("actual = %v, expected = -1", actual)
	}
}

/*
0 -> 1
1 -> 2
2 -> 3
3 -> 0
*/
func TestF(t *testing.T) {
	actual := closestMeetingNode([]int{1,2,3,0}, 1, 0)
	if actual != 1 {
		t.Fatalf("actual = %v, expected = 1", actual)
	}
}

func TestG(t *testing.T) {
	actual := closestMeetingNode([]int{5,4,5,4,3,6,-1}, 1, 0)
	if actual != -1 {
		t.Fatalf("actual = %v, expected = -1", actual)
	}
}
