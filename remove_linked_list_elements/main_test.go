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

func Test1263456(t *testing.T) {
	// 1->2->6->3->4->5->6
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}}

	actual := toIntSlice(removeElements(head, 6))
	expected := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func Test1123(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}

	actual := toIntSlice(removeElements(head, 1))
	expected := []int{2, 3}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func Test12344(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4}}}}}

	actual := toIntSlice(removeElements(head, 4))
	expected := []int{1, 2, 3}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}

func Test1(t *testing.T) {
	head := &ListNode{Val: 1}

	actual := toIntSlice(removeElements(head, 1))

	if len(actual) != 0 {
		t.Fatalf("%d", actual)
	}
}

func Test1221(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}

	actual := toIntSlice(removeElements(head, 2))
	expected := []int{1, 1}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%d", actual)
		}
	}
}
