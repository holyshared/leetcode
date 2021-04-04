package main

import "testing"

func TestA(t *testing.T) {
	cq := Constructor(3)
	cq.EnQueue(1)
	cq.EnQueue(2)
	cq.EnQueue(3)

	if cq.EnQueue(4) {
		t.Fatal("true")
	}
	if cq.Rear() != 3 {
		t.Fatalf("%d", cq.Rear())
	}
	if !cq.IsFull() {
		t.Fatal("false")
	}
	if !cq.DeQueue() {
		t.Fatal("false")
	}
	if !cq.EnQueue(4) {
		t.Fatal("false")
	}
	if cq.Rear() != 4 {
		t.Fatal("false")
	}
}
