package main

import (
	"testing"
)

func toIntSlice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}

	curr := head
	results := []int{}

	for curr.Next != nil {
		results = append(results, curr.Val)
		curr = curr.Next
	}
	results = append(results, curr.Val)
	return results
}

func Test12345(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	actual := toIntSlice(reverseList(head))
	expected := []int{5, 4, 3, 2, 1}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func Test12(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	actual := toIntSlice(reverseList(head))
	expected := []int{2, 1}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func Test1(t *testing.T) {
	head := &ListNode{Val: 1}
	actual := toIntSlice(reverseList(head))
	expected := []int{1}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func TestEmpty(t *testing.T) {
	actual := toIntSlice(reverseList(nil))

	if len(actual) != 0 {
		t.Fatalf("%d", actual)
	}
}
