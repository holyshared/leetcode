package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	orders := []*ListNode{}

	i := 0
	for curr.Next != nil {
		orders = append(orders, curr)
		curr = curr.Next
		i++
	}
	orders = append(orders, curr)

	rev, revHead := orders[i], orders[i]
	for j := len(orders) - 1; j >= 0; j-- {
		rev.Next = orders[j]
		rev = rev.Next
	}
	rev.Next = nil

	return revHead
}
