package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return nil
	}

	curr := head
	steps := map[*ListNode]bool{}

	for curr.Next != nil {
		_, ok := steps[curr]
		if ok {
			return curr
		}
		steps[curr] = true
		curr = curr.Next
	}

	return nil
}
