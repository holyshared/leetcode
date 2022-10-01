package main

func killProcess(pid []int, ppid []int, kill int) []int {
	graph := map[int][]int{}

	for i, id := range ppid {
		childs, has := graph[id]
		if has {
			childs = append(childs, pid[i])
			graph[id] = childs
		} else {
			graph[id] = []int{pid[i]}
		}
	}

	ans := []int{}
	queue := []int{kill}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		ans = append(ans, curr)

		for _, next := range graph[curr] {
			queue = append(queue, next)
		}
	}

	return ans
}