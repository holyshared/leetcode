package main

import "sort"

var N1 = 1 
var N2 = 2

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	visited := make([]int, len(edges))
	queue := [][]int{{node1, N1}, {node2, N2}}

	for len(queue) > 0 {
		size := len(queue)
		results := []int{}

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
	
			node, nnum := curr[0], curr[1]

			if visited[node] != 0 {
				if visited[node] != nnum {
					results = append(results, node)
				}
				continue
			}

			visited[node] = nnum

			if edges[node] == -1 {
				continue
			} else {
				queue = append(queue, []int{edges[node], nnum})
			}
		}

		if len(results) > 0 {
			sort.Slice(results, func (i, j int) bool { 
				return results[i] < results[j]
			})
			return results[0]
		}
	}

	return -1
}