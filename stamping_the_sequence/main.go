package main

import "math"

type Set struct {
	dist   map[int]int
	values []int
}

func NewSet() *Set {
	return &Set{
		dist:   map[int]int{},
		values: []int{},
	}
}

func (this *Set) Add(v int) {
	_, ok := this.dist[v]
	if !ok {
		this.dist[v] = len(this.dist)
	}
}

func (this *Set) Remove(v int) {
	delete(this.dist, v)
}

func (this *Set) Len() int {
	return len(this.dist)
}
func (this *Set) Contains(v int) bool {
	_, ok := this.dist[v]
	return ok
}
func (this *Set) IsEmpty() bool {
	return len(this.dist) <= 0
}
func (this *Set) Values() []int {
	values := make([]int, len(this.dist))

	for k, v := range this.dist {
		values[v] = k
	}
	return values
}

type Node struct {
	made *Set
	todo *Set
}

func NewNode(made, todo *Set) *Node {
	return &Node{
		made,
		todo,
	}
}

func movesToStamp(stamp string, target string) []int {
	M, N := len(stamp), len(target)

	queue := []int{}
	done := make([]bool, N)
	ans := []int{}
	A := []*Node{}

	tchars := []rune(target)
	schars := []rune(stamp)

	for i := 0; i <= N-M; i++ {
		made := NewSet()
		todo := NewSet()
		for j := 0; j < M; j++ {
			if tchars[i+j] == schars[j] {
				made.Add(i + j)
			} else {
				todo.Add(i + j)
			}
		}

		A = append(A, NewNode(made, todo))

		if todo.Len() <= 0 {
			ans = append(ans, i)
			for j := i; j < i+M; j++ {
				if !done[j] {
					queue = append(queue, j)
					done[j] = true
				}
			}
		}
	}

	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]

		for j := int(math.Max(0.0, float64(i-M+1))); j <= int(math.Min(float64(N-M), float64(i))); j++ {
			hasI := A[j].todo.Contains(i)

			if hasI {
				A[j].todo.Remove(i)

				if A[j].todo.IsEmpty() {
					ans = append(ans, j)
					made := A[j].made.Values()
					for _, m := range made {
						if !done[m] {
							queue = append(queue, m)
							done[m] = true
						}
					}
				}
			}
		}
	}

	for _, b := range done {
		if !b {
			return []int{}
		}
	}

	ret := make([]int, len(ans))
	t := 0
	for len(ans) > 0 {
		ret[t] = ans[len(ans)-1]
		ans = ans[:len(ans)-1]
		t++
	}

	return ret
}
