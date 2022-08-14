package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Solution struct {
	results [][]int
}

func (this *Solution) traverse(node *TreeNode, score int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		this.traverse(node.Left, score - 1)
	}

	this.results = append(this.results, []int{node.Val, score})

	if node.Right != nil {
		this.traverse(node.Right, score + 1)
	}
}

func (this *Solution) Grouping(node *TreeNode, score int) [][]int {
	this.traverse(node, 0)
	results := [][]int{[]int{this.results[0][0]}}

	j := 0
	curr := this.results[0][1]
	for i := 1; i < len(this.results); i++ {
		val, score := this.results[i][0], this.results[i][1]
		if curr == score {
			results[j] = append(results[j], val)
		} else {
			j++
			results = append(results, []int{})
			results[j] = append(results[j], val)
			curr = score
		}
	}
	return results
}

func verticalOrder(root *TreeNode) [][]int {
	sol := &Solution{ results: [][]int{} }
	return sol.Grouping(root, 0)
}