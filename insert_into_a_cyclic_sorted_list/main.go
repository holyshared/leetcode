package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	if aNode == nil {
		first := &Node{Val: x}
		first.Next = first
		return first
	}

	head := aNode

	if head == head.Next {
		head.Next = &Node{Val: x, Next: head}
		return head
	}

	prev := aNode
	curr := aNode.Next

	// 3 4 1
	for {
		if prev.Val <= x && x <= curr.Val {
			prev.Next = &Node{Val: x, Next: curr}
			return head
		} else if prev.Val > curr.Val {
			if x >= prev.Val || x <= curr.Val {
				prev.Next = &Node{Val: x, Next: curr}
				return head

			}
		}
		prev = curr
		curr = curr.Next

		if head == curr {
			break
		}
	}
	prev.Next = &Node{Val: x, Next: curr}

	return head
}
