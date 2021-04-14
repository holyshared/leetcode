package main

import (
	"testing"
)

func TestA(t *testing.T) {

	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
						}}}}}}

	actual := partition(l, 3)
	expected := []int{1, 2, 2, 4, 3, 5}

	i := 0
	for actual != nil {
		if actual.Val != expected[i] {
			t.Fatalf("%d:  %d != %d", i, actual.Val, expected[i])
		}
		actual = actual.Next
		i++
	}

}

func TestB(t *testing.T) {

	l := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
		},
	}

	actual := partition(l, 2)
	expected := []int{1, 2}

	i := 0
	for actual != nil {
		if actual.Val != expected[i] {
			t.Fatalf("%d:  %d != %d", i, actual.Val, expected[i])
		}
		actual = actual.Next
		i++
	}

}
