package main

import (
	"testing"
)

func TestA(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	expected := []int{4, 5, 1, 2, 3}
	actual := rotateRight(head, 2)

	for i := 0; actual != nil; i++ {
		if actual.Val != expected[i] {
			t.Fatalf("index[%d]: %d != %d", i, actual.Val, expected[i])
		}
		actual = actual.Next
	}
}

func TestB(t *testing.T) {
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}

	expected := []int{2, 0, 1}
	actual := rotateRight(head, 4)

	for i := 0; actual != nil; i++ {
		if actual.Val != expected[i] {
			t.Fatalf("index[%d]: %d != %d", i, actual.Val, expected[i])
		}
		actual = actual.Next
	}
}

func TestC(t *testing.T) {
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1}}

	expected := []int{0, 1}
	actual := rotateRight(head, 2)

	for i := 0; actual != nil; i++ {
		if actual.Val != expected[i] {
			t.Fatalf("index[%d]: %d != %d", i, actual.Val, expected[i])
		}
		actual = actual.Next
	}
}

func TestD(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	// 1 -> 2 -> 3
	// 1 3 -> 1 -> 2
	// 2 2 -> 3 -> 1
	// 3 1 -> 2 -> 3

	//   1 -> 2 -> 3 -> 4
	// 1 4 -> 1 -> 2 -> 3
	// 2 3 -> 4 -> 1 -> 2
	// 3 2 -> 3 -> 4 -> 1
	// 4 1 -> 2 -> 3 -> 4

	//	expected := []int{0, 1}
	actual := rotateRight(head, 2000000000)

	for actual != nil {
		actual = actual.Next
	}
	// 231

}
