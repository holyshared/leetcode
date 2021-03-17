package main

import "testing"

func Test1245367(t *testing.T) {
	root := &Node{
		Val:1,
		Left: &Node{
			Val:2,
			Left: &Node{
				Val:4,
			},
			Right: &Node{
				Val:5,
			},
		},
		Right: &Node{
			Val:3,
			Left: &Node{
				Val:6,
			},
			Right: &Node{
				Val:7,
			},
		},
	}
	actual := connect(root)

	l3 := actual.Left.Left
	actualL3 := []int{
		l3.Val,
		l3.Next.Val,
		l3.Next.Next.Val,
		l3.Next.Next.Next.Val,
	}

	expectL2 := []int{4,5,6,7}

	for i := 0; i < len(actualL3); i++ {
		if actualL3[i] != expectL2[i] {
			t.Fatalf("%d != %d", actualL3[i], expectL2[i])
		}
	}
}
