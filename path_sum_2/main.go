package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	target  int
	results [][]int
}

func (this *Solution) sum(node *TreeNode, paths []int, sum int) {
	if node == nil {
		return
	}

	if sum+node.Val == this.target && node.Left == nil && node.Right == nil {
		paths = append(paths, node.Val)
		out := make([]int, len(paths))
		copy(out, paths)
		this.results = append(this.results, out)
		return
	}
	paths = append(paths, node.Val)

	if node.Left != nil {
		this.sum(node.Left, paths, sum+node.Val)
	}
	if node.Right != nil {
		this.sum(node.Right, paths, sum+node.Val)
	}
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	sol := &Solution{target: targetSum, results: [][]int{}}
	sol.sum(root, []int{}, 0)
	return sol.results
}
