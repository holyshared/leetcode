package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	ll := 0

	curr := head

	for curr != nil {
		ll++
		curr = curr.Next
	}

	if (k-1 == ll-k) || ll <= 1 {
		return head
	}

	fcur := head
	for i := 0; i < k-1; i++ {
		fcur = fcur.Next
	}
	ecur := head
	for j := 0; j < ll-k; j++ {
		ecur = ecur.Next
	}

	fv := fcur.Val
	ev := ecur.Val
	fcur.Val = ev
	ecur.Val = fv

	return head
}
