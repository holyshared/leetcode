package main

import "testing"

func TestA(t *testing.T) {
	roots := &Node{
		Val: 1,
		Children: []*Node{
			&Node{Val: 3, Children: []*Node{
				&Node{Val: 5},
				&Node{Val: 6},
			}},
			&Node{Val: 2},
			&Node{Val: 4},
		},
	}
	actual := preorder(roots)
	expected := []int{1, 3, 5, 6, 2, 4}

	for i, val := range actual {
		if expected[i] != val {
			t.Fatalf("%d %d %d", i, val, expected[i])
		}
	}
}
