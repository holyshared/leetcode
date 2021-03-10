package main

import (
	"testing"
)

func Test341(t *testing.T) {
	last := &Node{Val: 1}
	head := &Node{Val: 3, Next: &Node{Val: 4, Next: last}}
	last.Next = head

	result := insert(head, 2)

	actual := []int{
		result.Val,                     // 3
		result.Next.Val,                // 4
		result.Next.Next.Val,           // 1
		result.Next.Next.Next.Val,      // 2
		result.Next.Next.Next.Next.Val, // 3
	}

	expected := []int{3, 4, 1, 2, 3}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("actual[%d] = %d", i, actual[i])
		}
	}
}

func TestEmpty(t *testing.T) {
	result := insert(nil, 1)
	if result.Val != 1 {
		t.Fatalf("actual = %d", result.Val)
	}
}

func Test1(t *testing.T) {
	head := &Node{Val: 1}
	head.Next = head

	result := insert(head, 0)

	actual := []int{
		result.Val,                // 1
		result.Next.Val,           // 0
		result.Next.Next.Val,      // 1
		result.Next.Next.Next.Val, // 0
	}

	expected := []int{1, 0, 1, 0}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("actual[%d] = %d", i, actual[i])
		}
	}
}

func Test333(t *testing.T) {
	last := &Node{Val: 3}
	head := &Node{Val: 3, Next: &Node{Val: 3}}
	head.Next = last
	last.Next = head

	result := insert(head, 0)

	actual := []int{
		result.Val,                // 0
		result.Next.Val,           // 3
		result.Next.Next.Val,      // 3
		result.Next.Next.Next.Val, // 3
	}

	expected := []int{3, 3, 0, 3}

	for i, v := range expected {
		if actual[i] != v {
			t.Fatalf("actual[%d] = %d", i, actual[i])
		}
	}
}
