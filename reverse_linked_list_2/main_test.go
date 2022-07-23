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
	expected := []int{1, 4, 3, 2, 5}

	actual := reverseBetween(root, 2, 4)
	actualList := nodeToList(actual)

	if !match(actualList, expected) {
		t.Fatalf("actual = %v, expected = %v", actualList, expected)
	}
}

func TestB(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := []int{2, 1, 3, 4, 5}

	actual := reverseBetween(root, 1, 2)
	actualList := nodeToList(actual)

	if !match(actualList, expected) {
		t.Fatalf("actual = %v, expected = %v", actualList, expected)
	}
}

func TestC(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := []int{1, 2, 3, 5, 4}

	actual := reverseBetween(root, 4, 5)
	actualList := nodeToList(actual)

	if !match(actualList, expected) {
		t.Fatalf("actual = %v, expected = %v", actualList, expected)
	}
}

func TestD(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	expected := []int{1, 2, 3, 4, 5}

	actual := reverseBetween(root, 1, 1)
	actualList := nodeToList(actual)

	if !match(actualList, expected) {
		t.Fatalf("actual = %v, expected = %v", actualList, expected)
	}
}
