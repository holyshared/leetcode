package main

func dfs(graph [][]int, curr int, routes []int, allRoutes [][]int) [][]int {
	if len(graph) - 1 == curr {
		routes = append(routes, curr)
		return append(allRoutes, routes)
	} else if len(graph[curr]) == 0 {
		return allRoutes
	}

	routes = append(routes, curr)

	for _, next := range graph[curr] {
		clonedRoutes := make([]int, len(routes))
		copy(clonedRoutes, routes)
		allRoutes = dfs(graph, next, clonedRoutes, allRoutes)
	}

	return allRoutes
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return dfs(graph, 0, []int{}, [][]int{})
}
