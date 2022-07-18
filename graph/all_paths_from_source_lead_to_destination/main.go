package main

var BLACK = 1
var GRAY = 2

func leadsToDest(graph [][]int, node int, dest int, states []int) bool {
	if states[node] != -1 {
		return states[node] == BLACK
	}

	if len(graph[node]) <= 0 {
		return node == dest
	}

	states[node] = GRAY

	// 次のルートから先が終端までいけた場合はOKにする
	for _, next := range graph[node] {
		if !leadsToDest(graph, next, dest, states) {
			return false
		}
	}

	states[node] = BLACK
	return true
}

func buildDigraph(n int, edges [][]int) [][]int {
	graph := make([][]int, n)

	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
	}

	return graph
}

func leadsToDestination(n int, edges [][]int, source int, destination int) bool {
	graph := buildDigraph(n, edges)
	states := make([]int, n)
	for i := 0; i < len(states); i++ {
		states[i] = -1
	}
	return leadsToDest(graph, source, destination, states)
}
