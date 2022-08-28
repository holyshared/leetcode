package main

func traverse(routes map[int][]int, ng []bool, pass []bool, curr int) int {
	if ng[curr] {
		return 0
	}
	if pass[curr] {
		return 0
	}

	pass[curr] = true

	nexts, has := routes[curr]
	if !has {
		return 1
	}

	sum := 0
	for _, next := range nexts {
		sum += traverse(routes, ng, pass, next)
	}
	return 1 + sum
}

func reachableNodes(n int, edges [][]int, restricted []int) int {
	routes := map[int][]int{}
	ng := make([]bool, n)
	pass := make([]bool, n)

	for i := 0; i < len(restricted); i++ {
		ng[restricted[i]] = true
	}

	for  _, edge := range edges {
		from, to := edge[0], edge[1]
		fromRoutes, ok := routes[from]
		if !ok {
			fromRoutes = []int{to}
		} else {
			fromRoutes = append(fromRoutes, to)
		}
		routes[from] = fromRoutes

		toRoutes, ok := routes[to]
		if !ok {
			toRoutes = []int{from}
		} else {
			toRoutes = append(toRoutes, from)
		}
		routes[to] = toRoutes
	}
	return traverse(routes, ng, pass, 0)
}