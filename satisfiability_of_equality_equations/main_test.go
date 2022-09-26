package main

import "testing"

func TestA(t *testing.T) {
	actual := equationsPossible([]string{"a==b","b!=a"})
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

func TestB(t *testing.T) {
	actual := equationsPossible([]string{"b!=a","a==b"})
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

func TestC(t *testing.T) {
	actual := equationsPossible([]string{"b==a","a==b"})
	if !actual {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}

func TestD(t *testing.T) {
	actual := equationsPossible([]string{"a!=a"})
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

func TestE(t *testing.T) {
	actual := equationsPossible([]string{"a!=b","b!=c","c!=a"})
	if !actual {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}

func TestF(t *testing.T) {
	actual := equationsPossible([]string{"a==b","b!=c","c==a"})
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}

// c:1、f:2、a:3、b:2、
func TestG(t *testing.T) {
	actual := equationsPossible([]string{"c==c","f!=a","f==b","b==c"})
	if !actual {
		t.Fatalf("actual = %v, expected = true", actual)
	}
}

// a b e c 
func TestH(t *testing.T) {
	actual := equationsPossible([]string{"a==b","e==c","b==c","a!=e",})
	if actual {
		t.Fatalf("actual = %v, expected = false", actual)
	}
}
