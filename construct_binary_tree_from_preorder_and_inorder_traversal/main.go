package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeBuilder struct {
	preOrders []int
	inOrders  []int
	preCursur int
	indexes   map[int]int
}

func NewTreeBuilder(preOrders []int, inOrders []int) *TreeBuilder {
	indexes := map[int]int{}
	for i, v := range inOrders {
		indexes[v] = i
	}

	return &TreeBuilder{
		preOrders: preOrders,
		inOrders:  inOrders,
		preCursur: 0,
		indexes:   indexes,
	}
}

// [3,9,20,15,7], inorder = [9,3,15,20,7]

func (this *TreeBuilder) r(left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := this.preOrders[this.preCursur]
	root := &TreeNode{Val: rootVal}

	index, _ := this.indexes[rootVal]

	this.preCursur++
	root.Left = this.r(left, index-1)
	root.Right = this.r(index+1, right)

	return root
}

func (this *TreeBuilder) Build() *TreeNode {
	return this.r(0, len(this.preOrders)-1)
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	builder := NewTreeBuilder(preorder, inorder)
	return builder.Build()
}
