package main

import "testing"

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
	removed := toIntSlice(removeNthFromEnd(head, 2))
	expected := []int{1, 2, 3, 5}

	for i := 0; i < len(expected); i++ {
		if removed[i] == expected[i] {
			continue
		}
		t.Fatalf("%d", removed)
	}
}

func Test1(t *testing.T) {
	head := &ListNode{Val: 1}
	removed := toIntSlice(removeNthFromEnd(head, 1))

	if len(removed) != 0 {
		t.Fatalf("%d", removed)
	}
}

func Test12(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	removed := toIntSlice(removeNthFromEnd(head, 1))

	expected := []int{1}

	for i := 0; i < len(expected); i++ {
		if removed[i] == expected[i] {
			continue
		}
		t.Fatalf("%d", removed)
	}
}

func Test123(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	removed := toIntSlice(removeNthFromEnd(head, 1))

	expected := []int{1, 2}

	for i := 0; i < len(expected); i++ {
		if removed[i] == expected[i] {
			continue
		}
		t.Fatalf("%d", removed)
	}
}

func Test123_3(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	removed := toIntSlice(removeNthFromEnd(head, 3))

	expected := []int{2, 3}

	for i := 0; i < len(expected); i++ {
		if removed[i] == expected[i] {
			continue
		}
		t.Fatalf("%d", removed)
	}
}
