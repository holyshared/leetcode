package main

func _climbStairs(n int, curr int, results []int) (int, []int) {
	if curr > n {
		return 0, results
	}
	if curr == n {
		return 1, results
	}
	if results[curr] > 0 {
		return results[curr], results
	}
	r1, results := _climbStairs(n, curr+1, results)
	r2, results := _climbStairs(n, curr+2, results)

	results[curr] = r1 + r2
	return results[curr], results
}

func climbStairs(n int) int {
	result, _ := _climbStairs(n, 0, make([]int, n+1))
	return result
}
