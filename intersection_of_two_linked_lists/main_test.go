package main

import "testing"

func Test41845_561845(t *testing.T) {
	st := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	at := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: st}}
	bt := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: st}}}

	actual := getIntersectionNode(at, bt)

	if actual == nil {
		t.Fatal("invalid point")
	}

	if actual.Val != 8 {
		t.Fatal("invalid point")
	}
}
