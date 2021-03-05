package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeSummary struct {
	Count int
	Total int
}

func traverse(node *TreeNode, levels map[int]*NodeSummary, l int) map[int]*NodeSummary {
	if node == nil {
		return levels
	}

	curr := levels

	val, ok := curr[l]
	if ok {
		val.Count++
		val.Total = val.Total + node.Val
		curr[l] = val
	} else {
		curr[l] = &NodeSummary{Total: node.Val, Count: 1}
	}

	if node.Left != nil {
		curr = traverse(node.Left, curr, l+1)
	}
	if node.Right != nil {
		curr = traverse(node.Right, curr, l+1)
	}
	return curr
}

func averageOfLevels(root *TreeNode) []float64 {
	results := []float64{}

	if root == nil {
		return results
	}

	levels := map[int]*NodeSummary{}
	levels[0] = &NodeSummary{Count: 1, Total: root.Val}

	levels = traverse(root.Left, levels, 1)
	levels = traverse(root.Right, levels, 1)

	for i := 0; i < len(levels); i++ {
		val, _ := levels[i]
		results = append(results, float64(val.Total)/float64(val.Count))
	}

	return results
}
