package main

type Node struct {
	Val int
	Next *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	curr := head
	copy := &Node{Val: curr.Val}
	copyHd := copy

	for curr != nil {
		if curr.Next != nil {
			copy.Next = &Node{Val: curr.Next.Val}
			copy = copy.Next
		} else {
			copy.Next = nil
		}
		curr = curr.Next
	}

	return copyHd
}