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
func _sortedArrayToBST(nums []int, left int, right int) *TreeNode {
	if (left > right) {
		return nil
	}
	mid := left + (right - left) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = _sortedArrayToBST(nums, left, mid - 1)
	root.Right = _sortedArrayToBST(nums, mid + 1, right)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	return _sortedArrayToBST(nums, 0, len(nums) - 1);
}

func main() {
	/*
	tree := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(tree.Val)
	fmt.Println(tree.Left.Val)
	fmt.Println(tree.Left.Left.Val)
	fmt.Println(tree.Right.Val)
	fmt.Println(tree.Right.Left.Val)
*/
	tree1 := sortedArrayToBST([]int{-1, 0, 1, 2})
	fmt.Println(tree1.Val)
	fmt.Println(tree1.Left.Val)
	fmt.Println(tree1.Right.Val)
	fmt.Println(tree1.Right.Right.Val)

	tree2 := sortedArrayToBST([]int{})
	fmt.Println(tree2)
}
