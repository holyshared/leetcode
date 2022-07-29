package main

import "testing"

func TestOneValueSlice(t *testing.T) {
	iter := Constructor([][]int{{1}, {2}})

	if !iter.HasNext() {
		t.Fatalf("actual = false, expected = true")
	}
	
	if actual := iter.Next(); actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}

	if !iter.HasNext() {
		t.Fatalf("actual = false, expected = true")
	}

	if actual := iter.Next(); actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}

	if iter.HasNext() {
		t.Fatalf("actual = true, expected = false")
	}

}


func TestTwoValueSlice(t *testing.T) {
	iter := Constructor([][]int{{1, 2}, {3, 4}})

	if actual := iter.Next(); actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
	if actual := iter.Next(); actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
	if actual := iter.Next(); actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
	if actual := iter.Next(); actual != 4 {
		t.Fatalf("actual = %d, expected = 4", actual)
	}
	if iter.HasNext() {
		t.Fatalf("actual = true, expected = false")
	}
}


func TestOneSlice(t *testing.T) {
	iter := Constructor([][]int{{1}})

	if !iter.HasNext() {
		t.Fatalf("actual = false, expected = true")
	}
	if actual := iter.Next(); actual != 1 {
		t.Fatalf("actual = %d, expected = 1", actual)
	}
	if iter.HasNext() {
		t.Fatalf("actual = true, expected = false")
	}
}

func TestEmptySlice(t *testing.T) {
	iter := Constructor([][]int{})

	if iter.HasNext() {
		t.Fatalf("actual = true, expected = false")
	}
}

func TestWithEmptySlice(t *testing.T) {
	iter := Constructor([][]int{{}, {3}})
	if !iter.HasNext() {
		t.Fatalf("actual = false, expected = true")
	}
	if actual := iter.Next(); actual != 3 {
		t.Fatalf("actual = %d, expected = 3", actual)
	}
	if iter.HasNext() {
		t.Fatalf("actual = true, expected = false")
	}
}
