package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _generateTrees(s, e int) []*TreeNode {
	allTrees := []*TreeNode{}
	if s > e {
		return append(allTrees, nil)
	}

	for i := s; i <= e; i++ {
		leftTrees := _generateTrees(s, i-1)
		rightTrees := _generateTrees(i+1, e)

		for _, l := range leftTrees {
			for _, r := range rightTrees {
				current := &TreeNode{Val: i}
				current.Left = l
				current.Right = r
				allTrees = append(allTrees, current)
			}
		}
	}
	return allTrees
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return _generateTrees(1, n)
}
