package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	var nodes = []*ListNode{}

	curr := head

	for curr != nil {
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	if (k-1 == len(nodes)-k) || len(nodes) <= 1 {
		return head
	}

	t1 := nodes[k-1]
	t2 := nodes[len(nodes)-k]
	t1v := t1.Val
	t2v := t2.Val
	t1.Val = t2v
	t2.Val = t1v

	return head
}
