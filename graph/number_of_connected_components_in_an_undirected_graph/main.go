package main

func find(representative []int, vertex int) int {
	if vertex == representative[vertex] {
		return vertex
	}
	representative[vertex] = find(representative, representative[vertex])

	return representative[vertex]
}

func combine(representative []int, size []int, vertex1 int, vertex2 int) int {
	vertex1 = find(representative, vertex1) // from
	vertex2 = find(representative, vertex2) // to

	if vertex1 == vertex2 {
		return 0
	} else {
		if size[vertex1] > size[vertex2] {
			size[vertex1] += size[vertex2]
			representative[vertex2] = vertex1
		} else {
			size[vertex2] += size[vertex1]
			representative[vertex1] = vertex2
		}
		return 1
	}
}

func countComponents(n int, edges [][]int) int {
	size := make([]int, n)
	representative := make([]int, n)

	for i := 0; i < n; i++ {
		representative[i] = i
		size[i] = 1
	}

	components := n
	for i := 0; i < len(edges); i++ {
		components -= combine(representative, size, edges[i][0], edges[i][1])
	}

	return components
}
