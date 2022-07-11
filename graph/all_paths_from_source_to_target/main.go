package main

func dfs(graph [][]int, curr int, routes []int, allRoutes [][]int) [][]int {
	if len(graph) - 1 == curr {
		return append(allRoutes, routes)
	}

	for _, next := range graph[curr] {
		clonedRoutes := make([]int, len(routes))
		copy(clonedRoutes, routes)
		clonedRoutes = append(clonedRoutes, next)
		allRoutes = dfs(graph, next, clonedRoutes, allRoutes)
	}

	return allRoutes
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return dfs(graph, 0, []int{0}, [][]int{})
}
