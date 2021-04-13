package main

import (
	"testing"
)

func TestA(t *testing.T) {

	lt := []*NestedInteger{
		&NestedInteger{
			Items: []*NestedInteger{
				&NestedInteger{Val: 1},
				&NestedInteger{Val: 1},
			},
		},
		&NestedInteger{Val: 2},
		&NestedInteger{
			Items: []*NestedInteger{
				&NestedInteger{Val: 1},
				&NestedInteger{Val: 1},
			},
		},
	}

	results := []int{}
	it := Constructor(lt)

	for it.HasNext() {
		results = append(results, it.Next())
	}

	expcted := []int{1, 1, 2, 1, 1}

	for i, v := range results {
		if v != expcted[i] {
			t.Fatalf("%d != %d", v, expcted[i])
		}
	}
}
