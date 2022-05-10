package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	val := root.Val
	closest := root.Val

	for root != nil {
		val = root.Val
		if math.Abs(float64(val)-target) < math.Abs(float64(closest)-target) {
			closest = val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closest
}
