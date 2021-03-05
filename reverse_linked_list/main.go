package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	prev *ListNode := nil

	// example: 1->2-3->4
	for curr != nil {
		nextTemp := curr.Next // 2 3 4
		curr.Next = prev // nil 1 2
		prev = curr // 1 2
		curr = nextTemp // 2 3
	}

	return prev
}
