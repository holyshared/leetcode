package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	pairs [][]int
}

func (this *Solution) calculateHeight(root *TreeNode) int {
	if root == nil {
		return -1
	}

	lh := this.calculateHeight(root.Left)
	rh := this.calculateHeight(root.Right)

	max := 0
	if lh < rh {
		max = rh
	} else {
		max = lh
	}

	ch := max + 1

	this.pairs = append(this.pairs, []int{ch, root.Val})

	return ch
}

func (this *Solution) resolve(root *TreeNode) [][]int {
	this.calculateHeight(root)

	// 深さで並び替える
	sort.Slice(this.pairs, func (i, j int) bool {
		return this.pairs[i][0] < this.pairs[j][0]
	})

	n, height, i := len(this.pairs), 0, 0

	results := [][]int{}

	for i < n {
		nums := []int{}
		for i < n && this.pairs[i][0] == height {
			nums = append(nums, this.pairs[i][1])
			i++
		}
		results = append(results, nums)
		height++
	}

	return results
}

func findLeaves(root *TreeNode) [][]int {
	sol := &Solution{pairs: [][]int{}}
	return sol.resolve(root)
}
