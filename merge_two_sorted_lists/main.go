package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 == nil && l2 == nil {
		return nil
	}

	h1 := l1
	h2 := l2
	var nh *ListNode = nil

	if h1.Val <= h2.Val {
		nh = h1
		h1 = h1.Next
	} else {
		nh = h2
		h2 = h2.Next
	}

	curr := nh

	for h1 != nil && h2 != nil {
		if h1.Val <= h2.Val {
			curr.Next = h1
			curr = h1
			h1 = h1.Next
		} else {
			curr.Next = h2
			curr = h2
			h2 = h2.Next
		}
	}

	if h1 != nil {
		curr.Next = h1
		curr = h1
	}
	if h2 != nil {
		curr.Next = h2
	}

	return nh
}
