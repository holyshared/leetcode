package main

import (
	"sort"
)

type UnionFind struct {
	root []int
	rank []int
}

func CreateUnionFind(size int) UnionFind {
	root := make([]int, size)
	rank := make([]int, size)

	for i := 0; i < size; i++ {
		root[i] = i
		rank[i] = 1
	}

	return UnionFind{
		root,
		rank,
	}
}

func (this *UnionFind) Find(x int) int {
	if x == this.root[x] {
		return x
	}
	this.root[x] = this.Find(this.root[x])
	return this.root[x]
}

func (this *UnionFind) Union(x int, y int) {
	rootX := this.Find(x)
	rootY := this.Find(y)

	if rootX != rootY {
		if this.rank[rootX] >= this.rank[rootY] {
			this.root[rootY] = rootX
			this.rank[rootX] += this.rank[rootY]
		} else {
			this.root[rootX] = rootY
			this.rank[rootY] += this.rank[rootX]
		}
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := CreateUnionFind(len(s))

	for _, edge := range pairs {
		uf.Union(edge[0], edge[1])
	}

	rootToComponent := map[int][]int{}

	for vertex := 0; vertex < len(s); vertex++ {
		root := uf.Find(vertex)
		v, ok := rootToComponent[root]
		if !ok {
			v = []int{}
		}
		rootToComponent[root] = append(v, vertex)
	}

	original := []rune(s)
	smallestString := []rune(s)

	for _, indices := range rootToComponent {
		characters := []rune{}
		for _, index := range indices {
			characters = append(characters, original[index])
		}
		sort.Slice(characters, func(i, j int) bool {
			return characters[i] < characters[j]
		})

		for index := 0; index < len(indices); index++ {
			smallestString[indices[index]] = characters[index]
		}
	}

	return string(smallestString)
}
