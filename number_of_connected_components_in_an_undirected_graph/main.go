package main

func dfs(adjList [][]int, visited []int, startNode int) {
	visited[startNode] = 1

	for i := 0; i < len(adjList[startNode]); i++ {
		if visited[adjList[startNode][i]] == 0 {
			dfs(adjList, visited, adjList[startNode][i])
		}
	}
}

func countComponents(n int, edges [][]int) int {
	components := 0
	visited := make([]int, n)
	adjList := make([][]int, n)

	for i := 0; i < n; i++ {
		adjList[i] = []int{}
	}

	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}

	for i := 0; i < n; i++ {
		if visited[i] == 0 {
			components++
			dfs(adjList, visited, i)
		}
	}
	return components
}
