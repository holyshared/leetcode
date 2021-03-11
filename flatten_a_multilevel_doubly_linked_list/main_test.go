package main

import (
	"fmt"
	"testing"
)

func Test123781112910456(t *testing.T) {
	// level1
	v1 := &Node{Val: 1}
	v2 := &Node{Val: 2}
	v3 := &Node{Val: 3}
	v4 := &Node{Val: 4}
	v5 := &Node{Val: 5}
	v6 := &Node{Val: 6}

	v1.Next = v2
	v2.Next = v3
	v3.Next = v4
	v4.Next = v5
	v5.Next = v6

	v6.Prev = v5
	v5.Prev = v4
	v4.Prev = v3
	v3.Prev = v2
	v2.Prev = v1

	// level2
	v7 := &Node{Val: 7}
	v8 := &Node{Val: 8}
	v9 := &Node{Val: 9}
	v10 := &Node{Val: 10}

	v7.Next = v8
	v8.Next = v9
	v9.Next = v10
	v10.Prev = v9
	v9.Prev = v8
	v8.Prev = v7

	// level 3
	v11 := &Node{Val: 11}
	v12 := &Node{Val: 12}
	v11.Next = v12
	v12.Prev = v11

	v3.Child = v7
	v8.Child = v11

	actual := flatten(v1)

	fmt.Println("actual next---------")
	next := actual
	for next.Next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
	fmt.Println(next.Val)

	fmt.Println("actual prev---------")
	prev := next
	for prev != nil {
		fmt.Println(prev.Val)
		prev = prev.Prev
	}

}

/*
 1 -> null
 2 -> null
 3 -> null
*/
// [1,null,2,null,3,null]

func Test123(t *testing.T) {

	v1 := &Node{Val: 1}
	v2 := &Node{Val: 2}
	v3 := &Node{Val: 3}
	v1.Child = v2
	v2.Child = v3

	actual := flatten(v1)

	fmt.Println("actual next---------")
	next := actual
	for next.Next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
	fmt.Println(next.Val)

	fmt.Println("actual prev---------")
	prev := next
	for prev != nil {
		fmt.Println(prev.Val)
		prev = prev.Prev
	}

}
