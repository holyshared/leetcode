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

	for curr != nil {
		if curr.Child != nil {
			tmp := curr.Next
			curr.Next = _flatten(curr.Child)
			last := _last(curr.Next)
			last.Next = tmp
			curr.Next.Prev = curr
			if tmp != nil {
				tmp.Prev = last
			}
			curr.Child = nil
		}
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	return root
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	return _flatten(root)
}
