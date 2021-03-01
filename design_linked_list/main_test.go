package main

import "testing"

func TestAddAtHead(t *testing.T) {
	l := Constructor()
	l.AddAtHead(1)
	l.AddAtHead(2)

	if l.Head.Val != 2 {
		t.Fatalf("fatal: %v", l.Head.Val)
	}
	if l.Head.Next.Val != 1 {
		t.Fatalf("fatal: %v", l.Head.Next.Val)
	}
}

func TestAddAtTail(t *testing.T) {
	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)

	if l.Head.Val != 1 {
		t.Fatalf("fatal: %v", l.Head.Val)
	}
	if l.Head.Next.Val != 2 {
		t.Fatalf("fatal: %v", l.Head.Next.Val)
	}
	if l.Head.Next.Next.Val != 3 {
		t.Fatalf("fatal: %v", l.Head.Next.Next.Val)
	}
}

func TestGet(t *testing.T) {
	var actual int

	l := Constructor()

	actual = l.Get(0)
	if actual != -1 {
		t.Fatalf("fatal: %v", actual)
	}

	l.AddAtTail(1)
	l.AddAtTail(2)
	l.AddAtTail(3)

	actual = l.Get(0)
	if actual != 1 {
		t.Fatalf("fatal: %v", actual)
	}
	actual = l.Get(1)
	if actual != 2 {
		t.Fatalf("fatal: %v", actual)
	}
	actual = l.Get(2)
	if actual != 3 {
		t.Fatalf("fatal: %v", actual)
	}

	actual = l.Get(3)
	if actual != -1 {
		t.Fatalf("fatal: %v", actual)
	}
}

func TestAddAtIndex(t *testing.T) {
	var actual int

	l := Constructor()
	l.AddAtTail(1)
	l.AddAtTail(2)

	// 1 3 2
	l.AddAtIndex(1, 3)
	actual = l.Get(0)
	if actual != 1 {
		t.Fatalf("fatal: %v", actual)
	}
	actual = l.Get(1)
	if actual != 3 {
		t.Fatalf("fatal: %v", actual)
	}
	actual = l.Get(2)
	if actual != 2 {
		t.Fatalf("fatal: %v", actual)
	}

	// 1 3 2
	// 4 1 3 2
	l.AddAtIndex(0, 4)

	actual = l.Get(0)
	if actual != 4 {
		t.Fatalf("fatal: %v", actual)
	}

}
