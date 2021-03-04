package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	acurr, bcurr := headA, headB

	for acurr != bcurr {
		if acurr == nil {
			acurr = headB
		} else {
			acurr = acurr.Next
		}

		if bcurr == nil {
			bcurr = headA
		} else {
			bcurr = bcurr.Next
		}
	}

	return acurr
}
