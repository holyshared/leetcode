package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	prev := head
	curr := head.Next

	for curr.Next != nil {
		temp := curr.Next
		curr.Next = prev

		prev = curr
		curr = temp
	}
	curr.Next = prev
	head.Next = nil

	return curr, head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	top := head
	start := head

	for i := 0; i < left-1; i++ {
		top = start        // 1
		start = start.Next // 2
	}

	end := start
	for i := 0; i < right-left; i++ {
		end = end.Next
	}
	bottom := end.Next

	end.Next = nil

	a, b := reverse(start)

	if left-1 > 0 {
		top.Next = a
		b.Next = bottom
		return head
	} else {
		top = a
		b.Next = bottom
		return top
	}

}
