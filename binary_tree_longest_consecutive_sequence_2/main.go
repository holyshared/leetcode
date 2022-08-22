package main

import "math"

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

func (this *Solution) longestPath(node *TreeNode) []int {
	if node == nil {
		return []int{0, 0}
	}

	inr := 1
	dcr := 1

	if node.Left != nil {
		left := this.longestPath(node.Left)
		if node.Val == node.Left.Val+1 {
			dcr = left[1] + 1
		} else if node.Val == node.Left.Val-1 {
			inr = left[0] + 1
		}
	}

	if node.Right != nil {
		right := this.longestPath(node.Right)
		if node.Val == node.Right.Val+1 {
			dcr = max(dcr, right[1]+1)
		} else if node.Val == node.Right.Val-1 {
			inr = max(inr, right[0]+1)
		}
	}

	this.max = max(this.max, dcr+inr-1)
	return []int{inr, dcr}
}

func longestConsecutive(root *TreeNode) int {
	sol := &Solution{max: math.MinInt32}
	sol.longestPath(root)
	return sol.max
}
