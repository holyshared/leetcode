package main

import "fmt"

type Node struct {
	Val int
	Next *Node
	Random *Node
}


func copyRandomLink(s *Node, dist map[*Node]*Node) *Node {
	if s == nil {
		return nil
	}

	if s.Random == nil {
		return nil
	}

	v, ok := dist[s.Random]
	if ok {
		return v
	} else {
		return nil
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head
	copy := &Node{Val: curr.Val}
	copyHd := copy

	dist := map[*Node]*Node{}
	dist[curr] = copy

	for curr != nil {
		if curr.Next != nil {
			copy.Next = &Node{Val: curr.Next.Val}
			copy = copy.Next
		} else {
			copy.Next = nil
		}
		curr = curr.Next
		dist[curr] = copy
	}

	curr = head
	copy = copyHd

	for curr != nil {
		if curr.Random != nil {
			copy.Random = copyRandomLink(curr, dist)
		}
		copy = copy.Next
		curr = curr.Next
	}

	return copyHd
}