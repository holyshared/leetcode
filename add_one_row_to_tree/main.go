package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _addOneRow(node *TreeNode, level int, v int, d int) {
	if node == nil {
		return
	}

	if level >= d-2 {
		if node.Left != nil {
			tmp := node.Left
			node.Left = &TreeNode{Val: v, Left: tmp}
		} else {
			node.Left = &TreeNode{Val: v}
		}
		if node.Right != nil {
			tmp := node.Right
			node.Right = &TreeNode{Val: v, Right: tmp}
		} else {
			node.Right = &TreeNode{Val: v}
		}

		return
	}

	if node.Left != nil {
		_addOneRow(node.Left, level+1, v, d)
	}
	if node.Right != nil {
		_addOneRow(node.Right, level+1, v, d)
	}
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if 0 > d-2 {
		return &TreeNode{Val: v, Left: root}
	}
	_addOneRow(root, 0, v, d)
	return root
}
