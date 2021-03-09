package main

import (
	"testing"
)

func Test243_564(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)
	expected := []int{7, 0, 8}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}

func Test243_554(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4}}}

	result := addTwoNumbers(l1, l2)
	expected := []int{7, 9, 7}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}

func Test099_01(t *testing.T) {
	l1 := &ListNode{Val: 0, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 0, Next: &ListNode{Val: 1}}

	result := addTwoNumbers(l1, l2)
	expected := []int{0, 0, 0, 1}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}

func Test0_0(t *testing.T) {
	l1 := &ListNode{Val: 0}
	l2 := &ListNode{Val: 0}

	result := addTwoNumbers(l1, l2)
	expected := []int{0}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}

func Test9999999_9999(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}

	result := addTwoNumbers(l1, l2)
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}

func Test9999_9999999(t *testing.T) {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}

	result := addTwoNumbers(l1, l2)
	expected := []int{8, 9, 9, 9, 0, 0, 0, 1}

	for i := 0; result != nil; i++ {
		if expected[i] != result.Val {
			t.Fatalf("%d != %d", expected[i], result.Val)
		}
		result = result.Next
	}
}
