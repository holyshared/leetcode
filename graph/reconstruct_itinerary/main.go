package main

import (
	"sort"
)

type Solution struct {
	flights int
	routes  map[string][]string
	visited map[string][]bool
	result  []string
}

func NewSolution(tickets [][]string) Solution {
	routes := map[string][]string{}
	visited := map[string][]bool{}

	starts := []string{}

	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		fromTo, ok := routes[from]
		if ok {
			routes[from] = append(fromTo, to)
		} else {
			routes[from] = []string{to}
			starts = append(starts, from)
		}
	}

	for _, start := range starts {
		toTickets, _ := routes[start]
		sort.Slice(toTickets, func(i, j int) bool {
			return toTickets[i] < toTickets[j]
		})
		routes[start] = toTickets
		visited[start] = make([]bool, len(toTickets))
	}

	return Solution{routes: routes, visited: visited, flights: len(tickets), result: []string{}}
}

func (this *Solution) Backtracking(curr string, used []string) bool {
	if len(used) == this.flights+1 {
		result := make([]string, len(used))
		copy(result, used)
		this.result = result
		return true
	}

	nexts, ok := this.routes[curr]
	if !ok {
		return false
	}

	i := 0
	marked := this.visited[curr]

	for _, next := range nexts {
		if !marked[i] {
			marked[i] = true
			used = append(used, next)
			ret := this.Backtracking(next, used)
			used = used[:len(used)-1]
			marked[i] = false

			if ret {
				return true
			}
		}
		i++
	}

	return false
}

func findItinerary(tickets [][]string) []string {
	solution := NewSolution(tickets)
	solution.Backtracking("JFK", []string{"JFK"})
	return solution.result
}
