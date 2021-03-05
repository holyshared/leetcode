package main

import (
	"testing"
)

func Test112233(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	nl := mergeTwoLists(l1, l2)

	i := 0
	expected := []int{1, 1, 2, 2, 3, 3}

	for nl != nil {
		if expected[i] != nl.Val {
			t.Fatal("failed")
		}
		nl = nl.Next
		i++
	}

}

func Test11223_1(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}

	nl := mergeTwoLists(l1, l2)

	i := 0
	expected := []int{1, 1, 2, 2, 3}

	for nl != nil {
		if expected[i] != nl.Val {
			t.Fatal("failed")
		}
		nl = nl.Next
		i++
	}

}

func Test11223_2(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	nl := mergeTwoLists(l1, l2)

	i := 0
	expected := []int{1, 1, 2, 2, 3}

	for nl != nil {
		if expected[i] != nl.Val {
			t.Fatal("failed")
		}
		nl = nl.Next
		i++
	}

}

func Test12345(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 2, Next: &ListNode{Val: 4}}

	nl := mergeTwoLists(l1, l2)

	i := 0
	expected := []int{1, 2, 3, 4, 5}

	for nl != nil {
		if expected[i] != nl.Val {
			t.Fatal("failed")
		}
		nl = nl.Next
		i++
	}
}
