package main

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	actual := wordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"})
	fmt.Println(actual)
}
