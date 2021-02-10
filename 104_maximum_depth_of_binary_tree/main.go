package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _maxDepth(root *TreeNode) int {
	var ldepth, rdepth int

	if root.Left == nil && root.Right == nil {
		return 0
	}

	if root.Left != nil {
		ldepth = _maxDepth(root.Left) + 1
	}

	if root.Right != nil {
		rdepth = _maxDepth(root.Right) + 1
	}

	if ldepth == rdepth {
		return ldepth
	}
	if ldepth > rdepth {
		return ldepth
	} else {
		return rdepth
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return _maxDepth(root) + 1
}

func main() {
	left := TreeNode{Val: 9}
	right := TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 17}}
	root := TreeNode{Val: 3, Left: &left, Right: &right}

	r := maxDepth(&root)
	fmt.Println(r)

	root2 := TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	r2 := maxDepth(&root2)
	fmt.Println(r2)

	var root3 *TreeNode
	root3 = nil
	r3 := maxDepth(root3)
	fmt.Println(r3)

	var root4 = TreeNode{Val: 1}
	r4 := maxDepth(&root4)
	fmt.Println(r4)
}
