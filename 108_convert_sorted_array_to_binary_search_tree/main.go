package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) <= 0 {
		return nil
	}

	if len(nums) == 1 {
		root = &TreeNode{Val: nums[0]}
		return root
	}

	if len(nums)%2 == 0 && len(nums) == 2 {
		root = &TreeNode{Val: nums[1], Left: &TreeNode{Val: nums[0]}}
		return root
	}
	mid := len(nums) / 2
	root = &TreeNode{Val: nums[mid]}
	left := sortedArrayToBST(nums[:mid])
	right := sortedArrayToBST(nums[(mid + 1):])
	root.Left = left
	root.Right = right
	return root
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

	tree := sortedArrayToBST([]int{-1, 0, 1, 2})
	fmt.Println(tree.Val)

	//	[-1,0,1,2]
}
