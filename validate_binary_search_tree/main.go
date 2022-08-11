package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
     5
   1    6 
      3   7
*/
func validate(node *TreeNode, low, high int) bool {
	if node == nil {
		return true
	}
	if node.Val <= low || node.Val >= high {
		return false
	}
	return validate(node.Right, node.Val, high) && validate(node.Left, low, node.Val)
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt, math.MaxInt)
}
