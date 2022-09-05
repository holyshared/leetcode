package main


func findOrder(numCourses int, prerequisites [][]int) []int {
	deps := make([][]int, numCourses)
	indegree := make([]int, numCourses) // 必要なコース数
	topologicalOrder := make([]int, numCourses)

    for _, course := range prerequisites {
		want, need := course[0], course[1]
		if len(deps[need]) <= 0 {
			deps[need] = []int{want}
		} else {
			deps[need] = append(deps[need], want)
		}
		indegree[want]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	i := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		topologicalOrder[i] = node
		i++

		if len(deps[node]) > 0 {
			for _, neighbor := range deps[node] {
				indegree[neighbor]--
				if indegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	if i == numCourses {
		return topologicalOrder
	}

	return []int{}
}