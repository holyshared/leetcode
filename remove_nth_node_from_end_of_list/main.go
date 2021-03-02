package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	nodes := map[int]*ListNode{}
	curr := head
	for curr.Next != nil {
		nodes[i] = curr
		i++
		curr = curr.Next
	}
	nodes[i] = curr

	if n == 1 && len(nodes) == 1 {
		return nil
	}

	// 1, 2, 3
	// 0, 1, 2
	var prev *ListNode
	if i-n < 0 {
		prev, _ = nodes[1]
		return prev
	} else {
		prev, _ = nodes[i-n]
		next, _ := nodes[i-n+2]
		prev.Next = next

		return head
	}
}
