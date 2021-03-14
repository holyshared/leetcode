package main

import (
	"fmt"
	"testing"
)

/*
        1
   2        7

3    4         8
  5    6     9


  123456789

	[1] [2 7] [3 4 8] [5 6 9]]
*/

func TestA(t *testing.T) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 9}}}}

	actual := levelOrder(root)

	fmt.Println(actual)

}
