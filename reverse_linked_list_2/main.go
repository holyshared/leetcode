package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	prev := head
	curr := head.Next

	for curr.Next != nil {
		temp := curr.Next
		curr.Next = prev

		prev = curr
		curr = temp
	}
	curr.Next = prev
	head.Next = nil

	fmt.Println("last--")
	fmt.Println(head)

	return curr, head
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	top := head
	start := head

	for i := 0; i < left - 1; i++ {
		top = start        // 1
		start = start.Next // 2
	}

	//curr := start
	end := start
	for i := 0; i < right - left; i++ {
		end = end.Next
	}
	bottom := end.Next

//	fmt.Println(start.Val)
//	fmt.Println(start.Next.Val)
//	fmt.Println(start.Next.Next.Val)

	end.Next = nil

	//fmt.Println(start.Next.Val)

	a, b := reverse(start)

	top.Next = a
	b.Next = bottom

//	fmt.Println(top)
	//fmt.Println(bottom)

///	a.Next = bottom
/*
	fmt.Println(top.Val)
	fmt.Println(start.Val)
	fmt.Println(start.Next.Val)

	fmt.Println(end.Val)
	fmt.Println(bottom.Val)
*/
	return top
}
