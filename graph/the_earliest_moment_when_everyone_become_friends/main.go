package main

import (
	"sort"
)

type UnionFind struct {
	group []int
	rank  []int
}

func CreateUnionFind(size int) UnionFind {
	group := make([]int, size)
	rank := make([]int, size)

	for person := 0; person < size; person++ {
		group[person] = person
		rank[person] = 0
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

func (this *UnionFind) Union(a, b int) bool {
	groupA := this.Find(a)
	groupB := this.Find(b)
	isMerged := false

	if groupA == groupB {
		return isMerged
	}

	isMerged = true

	if this.rank[groupA] > this.rank[groupB] {
		this.group[groupB] = groupA
	} else if this.rank[groupA] < this.rank[groupB] {
		this.group[groupA] = groupB
	} else {
		this.group[groupA] = groupB
		this.rank[groupB] += 1
	}

	return isMerged
}

func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		il, jl := logs[i], logs[j]
		return il[0] < jl[0]
	})

	groupCount := n
	uf := CreateUnionFind(n)

	for _, log := range logs {
		timestamp, friendA, friendB := log[0], log[1], log[2]

		if uf.Union(friendA, friendB) {
			groupCount -= 1
		}

		if groupCount == 1 {
			return timestamp
		}
	}

	return -1
}
