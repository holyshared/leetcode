package main

import (
	"testing"
)

/*
func TestA(t *testing.T) {
	c := Constructor()
	serialized := c.serialize(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 1},
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 5,
			},
		},
	})

	if serialized != "4(2(3)(1))(6(5))" {
		t.Fatalf("%s != %s", "4(2(3)(1))(6(5))", serialized)
	}
}

func TestB(t *testing.T) {
	c := Constructor()
	deserialized := c.deserialize("4(2(3)(1))(6(5)(7))")
	serialized := c.serialize(deserialized)

	if serialized != "4(2(3)(1))(6(5)(7))" {
		t.Fatalf("%s != %s", "4(2(3)(1))(6(5)(7))", serialized)
	}
}
*/
func TestC(t *testing.T) {
	c := Constructor()
	deserialized := c.deserialize("1,null,2,null,null,")
	serialized := c.serialize(deserialized)

	if serialized != "1,null,2,null,null," {
		t.Fatalf("%s != %s", "1,null,2,null,null,", serialized)
	}
}
