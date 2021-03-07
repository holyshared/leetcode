package main

import "testing"

func TestHashMap(t *testing.T) {
	actual := 0

	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	hm.Put(2, 1) // update the existing value

	actual = hm.Get(2) // returns 1
	if actual != 1 {
		t.Fatal("oops!!")
	}

	hm.Remove(2) // remove the mapping for 2

	actual = hm.Get(2)
	if actual != -1 {
		t.Fatal("oops!!")
	}
}
