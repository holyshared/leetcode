package main

/**
 *                     1
 *              2            2
 *           3     4       4      3
 *         5   6 7   8   8   7  6    5
 *
 */

// 2356478
// 2356478

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _left(node *TreeNode, ty int, values [][]int) [][]int {
	if node == nil {
		return values
	}

	values = append(values, []int{ty, node.Val})
	if node.Left != nil {
		values = _left(node.Left, 0, values)
	}
	if node.Right != nil {
		values = _left(node.Right, 1, values)
	}
	return values
}

func _right(node *TreeNode, ty int, values [][]int) [][]int {
	if node == nil {
		return values
	}
	values = append(values, []int{ty, node.Val})
	if node.Right != nil {
		values = _right(node.Right, 1, values)
	}
	if node.Left != nil {
		values = _right(node.Left, 0, values)
	}
	return values
}

func isSymmetric(node *TreeNode) bool {
	if node == nil {
		return false // ?
	}

	l := [][]int{}
	if node.Left != nil {
		l = _left(node.Left, 0, l)
	}

	r := [][]int{}
	if node.Right != nil {
		r = _right(node.Right, 1, r)
	}

	if len(l) != len(r) {
		return false
	}

	for i := 0; i < len(l); i++ {
		if l[i][0] == r[i][0] {
			return false
		}
		if l[i][1] != r[i][1] {
			return false
		}
	}

	return true
}
