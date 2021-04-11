package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	level := 0
	levels := [][]*TreeNode{}

	for len(queue) > 0 {
		levels = append(levels, []*TreeNode{})
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			levels[level] = append(levels[level], node)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
	}

	sum := 0
	for _, node := range levels[level-1] {
		sum += node.Val
	}
	return sum
}
