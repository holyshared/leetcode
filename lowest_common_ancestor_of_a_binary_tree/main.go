package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	ans *TreeNode
}

// 現在のノードから該当するノードを検索する
func (this *Solution) recurseTree(currentNode *TreeNode, p *TreeNode, q *TreeNode) bool {
	if currentNode == nil {
		return false
	}

	left, right := 0, 0
	if this.recurseTree(currentNode.Left, p, q) {
		left = 1
	}
	if this.recurseTree(currentNode.Right, p, q) {
		right = 1
	}

	mid := 0
	if currentNode == p || currentNode == q {
		mid = 1
	}

	// X 左右で見つかるパターン left = 1, right = 1, mid = 0
	// X 左で見つかり、現在のノードと一致 left = 1, right = 0, mid = 1
	// X 右で見つかり、現在のノードと一致 left = 0, right = 1, mid = 1
	//   左右で見つからず、現在のノードと一致 left = 0, right = 0, mid = 1
	if mid+left+right >= 2 {
		this.ans = currentNode
	}

	return (mid+left+right > 0)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	so := &Solution{
		ans: nil,
	}
	so.recurseTree(root, p, q)
	return so.ans
}
