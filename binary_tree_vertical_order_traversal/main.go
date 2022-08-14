package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	Column int
	Node   *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	output := [][]int{}

	if root == nil {
		return output
	}
	columnTable := map[int][]int{}

	column := 0
	queue := []*Pair{&Pair{Node: root, Column: column}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Node != nil {
			vals, has := columnTable[curr.Column]
			if !has {
				columnTable[curr.Column] = []int{curr.Node.Val}
			} else {
				vals = append(vals, curr.Node.Val)
				columnTable[curr.Column] = vals
			}
			queue = append(queue, &Pair{Node: curr.Node.Left, Column: curr.Column - 1})
			queue = append(queue, &Pair{Node: curr.Node.Right, Column: curr.Column + 1})
		}
	}

	keys := []int{}
	for k, _ := range columnTable {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		val, _ := columnTable[k]
		output = append(output, val)
	}

	return output
}
