package main

import (
	"testing"
)

func someSlice(actual, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}

	for i, val := range actual {
		if expected[i] != val {
			return false
		}
	}

	return true
}

func TestA(t *testing.T) {
	tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}

	expcted := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	acutal := findItinerary(tickets)

	if !someSlice(acutal, expcted) {
		t.Fatalf("acutal = %s, expected = %s", acutal, expcted)
	}
}

func TestB(t *testing.T) {
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}

	expcted := []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	acutal := findItinerary(tickets)

	if !someSlice(acutal, expcted) {
		t.Fatalf("acutal = %s, expected = %s", acutal, expcted)
	}
}


func TestC(t *testing.T) {
	tickets := [][]string{{"JFK","KUL"},{"JFK","NRT"},{"NRT","JFK"}}

	expcted := []string{"JFK","NRT","JFK","KUL"}
	acutal := findItinerary(tickets)

	if !someSlice(acutal, expcted) {
		t.Fatalf("acutal = %s, expected = %s", acutal, expcted)
	}
}
