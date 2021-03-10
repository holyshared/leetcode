package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	old := head
	var n int
	for n = 1; old.Next != nil; n++ {
		old = old.Next
	}
	old.Next = head

	new := head
	// 3 - 2000000000 % 3 - 1
	for i := 0; i < n-k%n-1; i++ {
		new = new.Next
	}
	newHd := new.Next
	new.Next = nil

	return newHd
}
