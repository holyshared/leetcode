package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getRoutes(node, t *TreeNode, routes []*TreeNode) ([]*TreeNode, bool) {
	if node == nil {
		return routes, false
	}

	if node == t {
		routes = append(routes, node)
		return routes, true
	}
	routes = append(routes, node)

	if node.Left != nil {
		routes, ok := getRoutes(node.Left, t, routes)
		if ok {
			return routes, ok
		}
		routes = routes[:(len(routes) - 1)]
	}

	if node.Right != nil {
		routes, ok := getRoutes(node.Right, t, routes)
		if ok {
			return routes, ok
		}
		routes = routes[:(len(routes) - 1)]
	}

	return routes, false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pRoutes, _ := getRoutes(root, p, []*TreeNode{})
	qRoutes, _ := getRoutes(root, q, []*TreeNode{})

	dep := map[*TreeNode]bool{}

	i := len(pRoutes) - 1
	j := len(qRoutes) - 1
	for i >= 0 || j >= 0 {
		if i >= 0 {
			pp := pRoutes[i]

			_, ok1 := dep[pp]
			if ok1 {
				return pp
			}
			dep[pp] = true
			i--
		}

		if j >= 0 {
			qp := qRoutes[j]

			_, ok2 := dep[qp]
			if ok2 {
				return qp
			}
			dep[qp] = true
			j--
		}
	}

	return nil
}
