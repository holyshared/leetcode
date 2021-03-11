package main

import "fmt"

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func _last(root *Node) *Node {
	curr := root
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}

func _flatten(root *Node) *Node {
	curr := root

	for curr.Next != nil {
		if curr.Child != nil {
			tmp := curr.Next
			curr.Next = _flatten(curr.Child)
			last := _last(curr.Next)
			last.Next = tmp

			fmt.Println("prev-----")
			fmt.Printf("%d -> %d\n", tmp.Val, last.Val)

			curr.Next.Prev = curr

			tmp.Prev = last
			curr.Child = nil
		}
		fmt.Println(curr.Val)
		curr = curr.Next
	}

	//	for curr.Prev != nil {
	//	curr = curr.Prev
	//	}
	fmt.Println("new head ---")
	fmt.Println(root.Val)

	return root
}

func flatten(root *Node) *Node {
	fmt.Println("start------")
	return _flatten(root)
}
