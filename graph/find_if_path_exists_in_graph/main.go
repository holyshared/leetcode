package main

func dfs(links [][]int, source int, visited []bool) {
	if visited[source] {
		return
	}
	visited[source] = true

	for i := 0; i < len(links[source]); i++ {
		edge := links[source][i]
		if !visited[edge] {
			dfs(links, edge, visited)
		}
	}
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	links := make([][]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		links[i] = []int{}
	}
	for _, edge := range edges {
		links[edge[0]] = append(links[edge[0]], edge[1])
		links[edge[1]] = append(links[edge[1]], edge[0])
	}
	dfs(links, source, visited)

	return visited[destination]
}
