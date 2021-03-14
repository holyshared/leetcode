package main

import "testing"

func Test12345_2(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := swapNodes(root, 2)

	actual := []int{
		result.Val,
		result.Next.Val,
		result.Next.Next.Val,
		result.Next.Next.Next.Val,
		result.Next.Next.Next.Next.Val,
	}
	expected := []int{1, 4, 3, 2, 5}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("index: %d - %d != %d", i, v, expected[i])
		}
	}
}

func Test12345_1(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := swapNodes(root, 1)

	actual := []int{
		result.Val,
		result.Next.Val,
		result.Next.Next.Val,
		result.Next.Next.Next.Val,
		result.Next.Next.Next.Next.Val,
	}
	expected := []int{5, 2, 3, 4, 1}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("index: %d - %d != %d", i, v, expected[i])
		}
	}
}

func Test12345_3(t *testing.T) {
	root := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := swapNodes(root, 3)

	actual := []int{
		result.Val,
		result.Next.Val,
		result.Next.Next.Val,
		result.Next.Next.Next.Val,
		result.Next.Next.Next.Next.Val,
	}
	expected := []int{1, 2, 3, 4, 5}

	for i, v := range actual {
		if v != expected[i] {
			t.Fatalf("index: %d - %d != %d", i, v, expected[i])
		}
	}
}
