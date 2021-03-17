package main

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	if  root == nil {
		return root
	}

	levels := [][]*Node{}
	queue := []*Node{root}
	level := 0

	for len(queue) > 0 {
		levels = append(levels, []*Node{})
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levels[level] = append(levels[level], node)

			if i - 1 >= 0 {
				levels[level][i - 1].Next = levels[level][i]
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			} 

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}

	return root
}