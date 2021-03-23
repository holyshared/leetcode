package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeBuilder struct {
	inOrders   []int
	postOrders []int
	postCursor int
	indexes    map[int]int
}

func NewBuilder(inOrders []int, postOrders []int) *TreeBuilder {
	indexes := map[int]int{}
	for i, v := range inOrders {
		indexes[v] = i
	}

	return &TreeBuilder{
		inOrders:   inOrders,
		postOrders: postOrders,
		postCursor: len(postOrders) - 1,
		indexes:    indexes,
	}
}

func (this *TreeBuilder) buildRecursive(left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := this.postOrders[this.postCursor]
	root := &TreeNode{Val: rootVal}

	index, _ := this.indexes[rootVal]

	this.postCursor--
	root.Right = this.buildRecursive(index+1, right)
	root.Left = this.buildRecursive(left, index-1)

	return root
}

func (this *TreeBuilder) Build() *TreeNode {
	return this.buildRecursive(0, len(this.inOrders)-1)
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	builder := NewBuilder(inorder, postorder)
	return builder.Build()
}
