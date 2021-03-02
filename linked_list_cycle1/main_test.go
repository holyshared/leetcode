package main

import "testing"

func Test3204(t *testing.T) {
	t1 := &ListNode{Val: 3}
	t2 := &ListNode{Val: 2}
	t3 := &ListNode{Val: 0}
	t4 := &ListNode{Val: -4}

	t1.Next = t2
	t2.Next = t3
	t3.Next = t4
	t4.Next = t2

	if hasCycle(t1) != true {
		t.Fatal("return false")
	}
}

func Test12(t *testing.T) {
	t1 := &ListNode{Val: 1}
	t2 := &ListNode{Val: 2}

	t1.Next = t2

	if hasCycle(t1) != false {
		t.Fatal("return true")
	}
}

func Test1(t *testing.T) {
	t1 := &ListNode{Val: 1}

	if hasCycle(t1) != false {
		t.Fatal("return true")
	}
}
