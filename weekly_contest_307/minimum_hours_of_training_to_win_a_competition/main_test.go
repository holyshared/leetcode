package main

import "testing"

func TestA(t *testing.T) {
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1,4,3,2}
	experience := []int{2,6,3,1}
	actual := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	if actual != 8 {
		t.Fatalf("actual = %d, expected = 8", actual)
	}
}
func TestB(t *testing.T) {
	initialEnergy := 2
	initialExperience := 4
	energy := []int{1}
	experience := []int{3}
	actual := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	if actual != 0 {
		t.Fatalf("actual = %d, expected = 0", actual)
	}
}

func TestC(t *testing.T) {
	initialEnergy := 2
	initialExperience := 4
	energy := []int{2}
	experience := []int{4}
	actual := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}

func TestD(t *testing.T) {
	initialEnergy := 1
	initialExperience := 1
	energy := []int{1,1,1,1}
	experience := []int{1,1,1,50}
	actual := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	if actual != 51 {
		t.Fatalf("actual = %d, expected = 51", actual)
	}
}

func TestE(t *testing.T) {
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1, 4}
	experience := []int{2, 5}
	actual := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
	if actual != 2 {
		t.Fatalf("actual = %d, expected = 2", actual)
	}
}
