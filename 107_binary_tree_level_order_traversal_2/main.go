package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func _levelOrderBottom(root *TreeNode, depth int, outputs [][]int) [][]int {
	if len(outputs) <= depth {
		outputs[depth] = []int{}
	}

	if root == nil {
		return outputs
	}
	outputs[depth] = append(outputs[depth], root.Val)

	if root.Left != nil {
		outputs = _levelOrderBottom(root.Left, depth + 1, outputs)
	}
	if root.Right != nil {
		outputs = _levelOrderBottom(root.Right, depth + 1, outputs)
	}
	return outputs
}




func levelOrderBottom(root *TreeNode) [][]int {
	results := [][]int{{}, {}, {}}
	reversedResults := [][]int{}

	results = _levelOrderBottom(root, 0, results)

	for i := len(results) - 1; i >= 0; i-- {
		reversedResults = append(reversedResults, results[i])
	}

	return reversedResults
}


func main() {
/*
	3
	/ \
 9  20
	 /  \
	15   7
*/

	left := TreeNode{Val: 9}
	right := TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}
	root := TreeNode{Val: 3, Left: &left, Right: &right}

	r := levelOrderBottom(&root)
	fmt.Println(r)
}