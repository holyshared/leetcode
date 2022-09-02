package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type Solution struct {
	max int
}

func (this *Solution) dfs(node *TreeNode, parent int, curr int) {
	if node == nil {
		return
	}
	if node.Val == parent+1 {
		curr++
	} else {
		curr = 1
	}
	this.max = max(curr, this.max)
	this.dfs(node.Left, node.Val, curr)
	this.dfs(node.Right, node.Val, curr)
}

func longestConsecutive(root *TreeNode) int {
	sol := &Solution{max: 0}
	sol.dfs(root, root.Val, 1)
	return sol.max
}
