package main

import "fmt"

func _climbStairs(n int, curr int, path []int, result [][]int) [][]int {
	fmt.Printf("%d > %d\n", curr, n)

	if curr >= n {
		dst := make([]int, len(path))
		copy(dst, path)
		result = append(result, dst)
		return result
	}

	if n >= curr+1 {
		path = append(path, 1)
		result = _climbStairs(n, curr+1, path, result)
		path = path[:(len(path) - 1)]
	}
	if n >= curr+2 {
		path = append(path, 2)
		result = _climbStairs(n, curr+2, path, result)
		path = path[:(len(path) - 1)]
	}
	return result
}

func climbStairs(n int) int {
	results := _climbStairs(n, 0, []int{}, [][]int{})
	fmt.Println(results)
	return len(results)
}
