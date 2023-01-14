package main

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	curr := head
	length := 0 
	for curr != nil {
		curr = curr.Next
		length++
	}
	length = length - n

	dummy := &ListNode{ Val: 0, Next: head }
	curr = dummy
	for length > 0 {
		curr = curr.Next
		length--
	}
	curr.Next = curr.Next.Next

	return dummy.Next
}