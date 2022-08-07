package main

import (
	"math"
	"sort"
)

type Solution struct {
	adj map[int]([][]int)
}

func (this *Solution) dfs(signalReceivedAt []int, currNode int, currTime int) {
	if currTime >= signalReceivedAt[currNode] {
		return
	}

	signalReceivedAt[currNode] = currTime

	if _, has := this.adj[currNode]; !has {
		return
	}

	edges, _ := this.adj[currNode]
	for _, edge := range edges {
		travelTime, neighborNode := edge[0], edge[1]
		this.dfs(signalReceivedAt, neighborNode, currTime+travelTime)
	}
}

func (this *Solution) networkDelayTime(times [][]int, n int, k int) int {
	for _, time := range times {
		source := time[0]
		dest := time[1]
		travelTime := time[2]

		routes, has := this.adj[source]
		if !has {
			this.adj[source] = [][]int{}
			routes = this.adj[source]
		}
		routes = append(routes, []int{travelTime, dest})
		this.adj[source] = routes
	}

	for source, routes := range this.adj {
		sort.Slice(routes, func(i, j int) bool {
			return routes[i][0] < routes[j][0]
		})
		this.adj[source] = routes
	}

	signalReceivedAt := make([]int, n+1)
	for i := 0; i < len(signalReceivedAt); i++ {
		signalReceivedAt[i] = math.MaxInt
	}

	this.dfs(signalReceivedAt, k, 0)

	answer := math.MinInt
	for node := 1; node <= n; node++ {
		if answer < signalReceivedAt[node] {
			answer = signalReceivedAt[node]
		}
	}

	if answer == math.MaxInt {
		return -1
	} else {
		return answer
	}
}

func networkDelayTime(times [][]int, n int, k int) int {
	sol := &Solution{adj: map[int][][]int{}}
	return sol.networkDelayTime(times, n, k)
}
