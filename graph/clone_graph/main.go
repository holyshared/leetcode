package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func _cloneGraph(node *Node, memo map[int]*Node) *Node {
	cloneNode, ok := memo[node.Val]

	if ok {
		return cloneNode
	}

	clone := &Node{Val: node.Val, Neighbors: []*Node{}}

	memo[node.Val] = clone

	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, _cloneGraph(neighbor, memo))
	}


	return clone 
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	memo := map[int]*Node{}
	return _cloneGraph(node, memo)
}
