package main

import (
	"testing"
)

func nodeToList(node *ListNode) []int {
	curr := node
	results := []int{}

	for curr.Next != nil {
		results = append(results, curr.Val)
		curr = curr.Next
	}
	results = append(results, curr.Val)

	return results
}

func match(actual, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, val := range actual {
		if val != expected[i] {
			return false
		}
	}

	return true
}


func TestA(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := []int{1, 4,3,2, 5}

	actual := reverseBetween(root, 2, 4)
	actualList := nodeToList(actual)

	if !match(actualList, expected) {
		t.Fatalf("actual = %v, expected = %v", actualList, expected)
	}
}
