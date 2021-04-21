package main

type Node struct {
	Val      int
	Children []*Node
}

func _preorder(root *Node, results []int) []int {
	results = append(results, root.Val)

	for _, node := range root.Children {
		results = _preorder(node, results)
	}

	return results
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	results := _preorder(root, []int{})
	return results
}
