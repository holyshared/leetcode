package main

import (
  "sort"
)

func moveRight(people [][]int, s int) [][]int {
	for i := len(people) - 1; i > s; i-- {
		people[i] = people[i - 1]
	}
	people[s] = []int{}
	return people
}

func reconstructQueue(people [][]int) [][]int {
  sort.Slice(people, func (i, j int) bool {
    p1, p2 := people[i], people[j]
    if p1[0] == p2[0] {
	  	return p1[1] < p2[1]
		} else {
		  return p1[0] > p2[0]
		}
  })

	queue := make([][]int, len(people))

	for i := 0; i < len(people); i++ {
		pos := people[i][1]
		if len(queue[pos]) > 0 {
			queue = moveRight(queue, pos)
		}
		queue[people[i][1]] = people[i]
	}
  return queue
}


