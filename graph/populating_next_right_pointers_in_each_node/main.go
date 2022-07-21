package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		levels := []*Node{}

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			levels = append(levels, curr)

			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		for i := 1; i < len(levels); i++ {
			levels[i - 1].Next = levels[i]
		}
	}

	return root
}
