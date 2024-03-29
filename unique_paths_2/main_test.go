package main

import "testing"

func TestA(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	actual := uniquePathsWithObstacles(obstacleGrid)
	if actual != 2 {
		t.Fatalf("%d", actual)
	}
}

func TestB(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 1},
		{0, 0},
	}
	actual := uniquePathsWithObstacles(obstacleGrid)
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}

func TestC(t *testing.T) {
	obstacleGrid := [][]int{
		{1},
	}
	actual := uniquePathsWithObstacles(obstacleGrid)
	if actual != 0 {
		t.Fatalf("%d", actual)
	}
}

func TestD(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0},
	}
	actual := uniquePathsWithObstacles(obstacleGrid)
	if actual != 1 {
		t.Fatalf("%d", actual)
	}
}
