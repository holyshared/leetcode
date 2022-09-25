package main

import "testing"

func TestA(t *testing.T) {
	actual := sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})
	t.Fatalf("actual = %v", actual)
}
