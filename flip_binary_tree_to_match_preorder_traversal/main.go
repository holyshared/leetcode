package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	flipped []int
	index   int
	voyage  []int
}

func NewSolution(voyage []int) *Solution {
	return &Solution{
		flipped: []int{},
		index:   0,
		voyage:  voyage,
	}
}

func (this *Solution) flipMatchVoyage(node *TreeNode) []int {
	this.dfs(node)
	if len(this.flipped) > 0 && this.flipped[0] == -1 {
		this.flipped = []int{-1}
	}
	return this.flipped
}

func (this *Solution) dfs(node *TreeNode) {
	if node != nil {
		if node.Val != this.voyage[this.index] {
			this.flipped = []int{-1}
			return
		}
		this.index++

		if this.index < len(this.voyage) && node.Left != nil && node.Left.Val != this.voyage[this.index] {
			this.flipped = append(this.flipped, node.Val)
			this.dfs(node.Right)
			this.dfs(node.Left)
		} else {
			this.dfs(node.Left)
			this.dfs(node.Right)
		}
	}
}

// [1,2,3], voyage = [1,3,2]
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	s := NewSolution(voyage)
	return s.flipMatchVoyage(root)
}
