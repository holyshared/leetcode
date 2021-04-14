package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1,4,3,2,5,2], x = 3
func partition(head *ListNode, x int) *ListNode {
	tmp1 := &ListNode{Val: 0}
	tmp2 := &ListNode{Val: 0}

	low := tmp1
	high := tmp2

	curr := head

	for curr != nil {
		tmp := curr.Next
		if curr.Val < x {
			low.Next = curr
			low = low.Next
		} else {
			high.Next = curr
			high = high.Next
		}
		curr = tmp
	}

	low.Next = tmp2.Next
	high.Next = nil

	return tmp1.Next
}
