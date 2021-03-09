package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := false
	c1, c2 := l1, l2
	var hd *ListNode = nil
	curr := hd

	for c1 != nil && c2 != nil {
		sum := c1.Val + c2.Val
		if carry == true {
			sum++
			carry = false
		}
		if sum >= 10 {
			r := sum - 10
			if curr != nil {
				curr.Next = &ListNode{Val: r}
				curr = curr.Next
			} else {
				curr = &ListNode{Val: r}
				hd = curr
			}
			carry = true
		} else {
			if curr != nil {
				curr.Next = &ListNode{Val: sum}
				curr = curr.Next
			} else {
				curr = &ListNode{Val: sum}
				hd = curr
			}
			carry = false
		}
		c1 = c1.Next
		c2 = c2.Next
	}

	if c1 != nil && c2 == nil {
		for c1 != nil {
			r := c1.Val
			if carry == true {
				r++
			}
			if r >= 10 {
				curr.Next = &ListNode{Val: 0}
				carry = true
			} else {
				curr.Next = &ListNode{Val: r}
				carry = false
			}
			curr = curr.Next
			c1 = c1.Next
		}
	} else if c1 == nil && c2 != nil {
		for c2 != nil {
			r := c2.Val
			if carry == true {
				r++
			}
			if r >= 10 {
				curr.Next = &ListNode{Val: 0}
				carry = true
			} else {
				curr.Next = &ListNode{Val: r}
				carry = false
			}
			curr = curr.Next
			c2 = c2.Next
		}
	}

	if carry == true {
		curr.Next = &ListNode{Val: 1}
	}

	return hd
}
