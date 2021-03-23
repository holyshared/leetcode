package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {

	actual := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(actual)

}
