package main

import (
	"fmt"
	"testing"
)

// [3,9,20,15,7], inorder = [9,3,15,20,7]

func TestA(t *testing.T) {

	actual := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(actual)

}
