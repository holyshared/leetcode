package main

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}
	return find(parent, parent[i])
}

func union(parent []int, x int, y int) []int {
	xset := find(parent, x)
	yset := find(parent, y)
	if xset != yset {
		parent[xset] = yset
	}
	return parent
}

// union find
func findCircleNum(isConnected [][]int) int {
	parent := make([]int, len(isConnected))

	for i := 0; i < len(isConnected); i++ {
		parent[i] = -1
	}

	for i := 0; i < len(isConnected); i++ {
		for j := 0; j < len(isConnected); j++ {
			if isConnected[i][j] == 1 && i != j {
				parent = union(parent, i, j)
			}
		}
	}
	count := 0
	for i := 0; i < len(parent); i++ {
		if parent[i] == -1 {
			count++
		}
	}

	return count
}
