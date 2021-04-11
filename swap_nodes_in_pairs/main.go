package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr := head      // 1
	next := head.Next // 2

	curr.Next = swapPairs(next.Next)
	next.Next = curr

	return next
}
