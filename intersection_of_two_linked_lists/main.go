package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	steps := map[*ListNode]bool{}
	acurr, bcurr := headA, headB
	ok := false

	for acurr != nil || bcurr != nil {
		if acurr != nil {
			_, ok = steps[acurr]
			if ok {
				return acurr
			}
			steps[acurr] = true
			acurr = acurr.Next
		}
		if bcurr != nil {
			_, ok = steps[bcurr]
			if ok {
				return bcurr
			}
			steps[bcurr] = true
			bcurr = bcurr.Next
		}
	}

	return nil
}
