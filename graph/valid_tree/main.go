package main

func validTree(n int, edges [][]int) bool {
	adjacencyList := [][]int{}
	for i := 0; i < n; i++ {
		adjacencyList = append(adjacencyList, []int{})
	}

	// From -> To、To -> Fromの隣接する情報を生成する
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}

	parent := map[int]int{}
	parent[0] = -1

	stack := []int{0}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, neighbour := range adjacencyList[node] {
			if parent[node] == neighbour {
				continue
			}
			// 最初のノードに既に親がいる場合はinvalid
			_, containsKey := parent[neighbour]
			if containsKey {
				return false
			}
			// キューに隣接するノードを積み、隣接するノードの親に自信を追加する
			stack = append(stack, neighbour)
			parent[neighbour] = node
		}
	}

	return len(parent) == n
}
