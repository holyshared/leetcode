package main

import (
	"sort"
)

type UnionFind struct {
	group []int
	rank  []int
}

func CreateUnionFind(size int) UnionFind {
	group := make([]int, size+1)
	rank := make([]int, size+1)
	for i := 0; i < size+1; i++ {
		group[i] = i
		rank[i] = 0
	}

	return UnionFind{
		group,
		rank,
	}
}

func (this *UnionFind) Find(person int) int {
	if this.group[person] != person {
		this.group[person] = this.Find(this.group[person])
	}
	return this.group[person]
}

func (this *UnionFind) Union(person1, person2 int) bool {
	group1 := this.Find(person1)
	group2 := this.Find(person2)

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

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	orderedEdges := [][]int{}

	// 仮想の家0を置き、井戸のコストをパイプのフォーマットと同じにする
	// 0 to i = cost wells[i]
	for i := 0; i < len(wells); i++ {
		orderedEdges = append(orderedEdges, []int{0, i + 1, wells[i]})
	}

	// パイプのコストを加える
	for i := 0; i < len(pipes); i++ {
		edge := pipes[i]
		orderedEdges = append(orderedEdges, edge)
	}

	// コストが少ない順に並び替える
	sort.Slice(orderedEdges, func(i, j int) bool {
		a, b := orderedEdges[i], orderedEdges[j]
		return a[2] < b[2]
	})

	uf := CreateUnionFind(n)

	// ノードを接続できたら、コストを合算していく
	// すでに接続されている場合は除外する
	totalCost := 0
	for _, edge := range orderedEdges {
		house1 := edge[0]
		house2 := edge[1]
		cost := edge[2]
		if uf.Union(house1, house2) {
			totalCost += cost
		}
	}

	return totalCost
}
