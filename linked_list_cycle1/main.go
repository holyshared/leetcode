package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return false
	}

	curr := head
	steps := map[*ListNode]bool{}
	pos := 0

	for curr.Next != nil {
		_, ok := steps[curr]
		if ok {
			return true
		}
		steps[curr] = true
		pos++
		curr = curr.Next
	}

	return false
}
