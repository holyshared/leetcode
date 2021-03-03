package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head

	for newHead != nil && newHead.Val == val {
		newHead = newHead.Next
	}
	curr := newHead

	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			// 1 2a 2b 1
			// 1 2b 1
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return newHead
}
