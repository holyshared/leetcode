package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	level := 0
	levels := [][]int{}
	queue := []*Node{root}

	for len(queue) > 0 {
		size := len(queue)
		levels = append(levels, []int{})

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			levels[level] = append(levels[level], curr.Val)

			for _, child := range curr.Children {
				queue = append(queue, child)
			}
		}
		level++
	}

	return levels
}
