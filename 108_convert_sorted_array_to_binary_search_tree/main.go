package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0 + (5 - 0) -10, -3, 0, 5, 9
// 0 + (2 - 0) -10, -3
func buildToBST(nums []int, left, right int) *TreeNode {
	if (left > right) {
		return nil
	}
	mid := (left + right) / 2

	if (left + right) % 2 == 1 {
		mid++
	}

	node := &TreeNode{Val: nums[mid],}
	node.Left = buildToBST(nums, left, mid - 1)
	node.Right = buildToBST(nums, mid + 1, right)
	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	return buildToBST(nums, 0, len(nums) - 1)
}