package main

import "testing"

// [[7,null],[13,0],[11,4],[10,2],[1,0]]
func TestA(t *testing.T) {
	v7 := &Node{Val:7}
	v13 := &Node{Val:13}
	v11 := &Node{Val:11}
	v10 := &Node{Val:10}
	v1 := &Node{Val:1}

	v7.Next = v13
	v13.Next = v11
	v11.Next = v10
	v10.Next = v1

	v7.Random = nil
	v13.Random = v7
	v11.Random = v1
	v10.Random = v11
	v1.Random = v7

	acutal := copyRandomList(v7)

	e7 := acutal
	e13 := acutal.Next
	e11 := acutal.Next.Next
	e10 := acutal.Next.Next.Next
	e1 := acutal.Next.Next.Next.Next

	if e7.Next != e13 {
		t.Fatal("oops!!")
	}
	if e13.Next != e11 {
		t.Fatal("oops!!")
	}
	if e11.Next != e10 {
		t.Fatal("oops!!")
	}
	if e10.Next != e1 {
		t.Fatal("oops!!")
	}
	if e1.Next != nil {
		t.Fatal("oops!!")
	}





	if e7.Random != nil {
		t.Fatal("oops!!")
	}
	if e13.Random != e7 {
		t.Fatal("oops!!")
	}
	if e11.Random != e10 {
		t.Fatal("oops!!")
	}
	if e10.Random != e11 {
		t.Fatal("oops!!")
	}
	if e1.Random != e7 {
		t.Fatal("oops!!")
	}
}