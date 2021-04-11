package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	actual := swapPairs(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4},
			},
		},
	})
	expexted := []int{2, 1, 4, 3}
	curr := actual

	for _, v := range expexted {
		fmt.Println(curr.Val)
		if curr.Val != v {
			t.Fatalf("%d != %d", curr.Val, v)
		}
		curr = curr.Next
	}

}
