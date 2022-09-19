package main

import "github.com/emirpasic/gods/sets/hashset"

func findMinHeightTrees(n int, edges [][]int) []int {
	if n < 2 {
		centroids := []int{}
		for  i := 0; i < n; i++ {
			centroids = append(centroids, i)
		}
		return centroids
	}

	neighbors := []*hashset.Set{}
	for i := 0; i < n; i++ {
		neighbors = append(neighbors, hashset.New())
	}

	for _, edge := range edges {
		start, end := edge[0], edge[1]
		neighbors[start].Add(end)
		neighbors[end].Add(start)
	}

	leaves := []int{}

	for i := 0; i < n; i++ {
		if neighbors[i].Size() == 1 {
			leaves = append(leaves, i)
		}
	}

	remainingNodes := n
	for remainingNodes > 2 {
		newLeaves := []int{}
		remainingNodes -= len(leaves)

		for _, leaf := range leaves {
			neighbor := neighbors[leaf].Values()[0].(int)
			neighbors[neighbor].Remove(leaf)

			if neighbors[neighbor].Size() == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
		}

		leaves = newLeaves
	}

	return leaves
}