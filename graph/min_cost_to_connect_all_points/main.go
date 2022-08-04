package main

import (
	"math"
	"sort"
)

type UnionFind struct {
	group []int
	rank  []int
}

func CreateUnionFind(size int) UnionFind {
	group := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		group[i] = i
	}

	return UnionFind{
		group: group, rank: rank}

}

func (this *UnionFind) Find(node int) int {
	if this.group[node] != node {
		this.group[node] = this.Find(this.group[node])
	}
	return this.group[node]
}

func (this *UnionFind) Union(node1, node2 int) bool {
	group1 := this.Find(node1)
	group2 := this.Find(node2)

	if group1 == group2 {
		return false
	}

	if this.rank[group1] > this.rank[group2] {
		this.group[group2] = group1
	} else if this.rank[group1] < this.rank[group2] {
		this.group[group1] = group2
	} else {
		this.group[group1] = group2
		this.rank[group2] += 1
	}

	return true
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	allEdges := [][]int{}

	for currNext := 0; currNext < n; currNext++ {
		for nextNext := currNext + 1; nextNext < n; nextNext++ {
			x := int(math.Abs(float64(points[currNext][0] - points[nextNext][0])))
			y := int(math.Abs(float64(points[currNext][1] - points[nextNext][1])))

			weight := x + y
			allEdges = append(allEdges, []int{weight, currNext, nextNext})
		}
	}

	// weightでソートする
	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i][0] < allEdges[j][0]
	})

	uf := CreateUnionFind(n)
	cost := 0
	edgesUsed := 0

	for i := 0; i < len(allEdges) && edgesUsed < n-1; i++ {
		edge := allEdges[i]
		weight, node1, node2 := edge[0], edge[1], edge[2]

		if uf.Union(node1, node2) {
			cost += weight
			edgesUsed++
		}
	}

	return cost
}
