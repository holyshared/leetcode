package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left, right := root.Left, root.Right
	root.Left = nil

	last := root

	if left != nil {
		next := _flatten(left)
		if next != nil {
			last.Right = next
		} else {
			last.Right = left
		}
		for last.Right != nil {
			last = last.Right
		}
	}

	if right != nil {
		next := _flatten(right)
		if next != nil {
			last.Right = next
		} else {
			last.Right = right
		}
		for last.Right != nil {
			last = last.Right
		}
	}

	return root
}

func flatten(root *TreeNode) {
	_flatten(root)
}
