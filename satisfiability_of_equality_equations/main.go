package main

func dfs(node int, c int, color []int, graph [][]int) []int {
	if color[node] == -1 {
		color[node] = c
		for _, nei := range graph[node] {
			color = dfs(nei, c, color, graph)
		}
	}
	return color
}

func equationsPossible(equations []string) bool {
	graph := make([][]int, 26)

	for i, _ := range graph {
		graph[i] = []int{}
	}
	not := []string{}

	for _, equation := range equations {
		egn := []rune(equation)
		if egn[1] == '=' {
			a, b := egn[0], egn[3]
			x := a - 'a'
			y := b - 'a'
			graph[x] = append(graph[x], int(y))
			graph[y] = append(graph[y], int(x))
		} else {
			not = append(not, equation)
		}
	}

	color := make([]int, 26)

	for i, _ := range color {
		color[i] = -1
	}

	for i := 0; i < 26; i++ {
		if color[i] == -1 {
			color = dfs(i, i, color, graph)
		}
	}

	for _, equation := range not {
		egn := []rune(equation)
		a, b := egn[0], egn[3]
		x := a - 'a'
		y := b - 'a'
		if color[x] == color[y] {
			return false
		}
	}
	return true
}
