package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func _str2tree(nodes []string) (*TreeNode, []string) {
	if len(nodes) <= 0 {
		return nil, nodes
	}

	if nodes[0] == "null" {
		nodes = nodes[1:]
		return nil, nodes
	}
	v, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{Val: v}
	nodes = nodes[1:]

	l, nodes := _str2tree(nodes)
	root.Left = l

	r, nodes := _str2tree(nodes)
	root.Right = r

	return root, nodes
}

func str2tree(s string) *TreeNode {
	root, _ := _str2tree(strings.Split(s, ","))
	return root
}

func _tree2str(node *TreeNode, s string) string {
	if node == nil {
		s += "null,"
	} else {
		s += strconv.Itoa(node.Val) + ","
		s = _tree2str(node.Left, s)
		s = _tree2str(node.Right, s)
	}
	return s
}

func tree2str(root *TreeNode) string {
	return _tree2str(root, "")
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return tree2str(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return str2tree(data)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
