package main

func backtrack(remain int, k int, comb []int, nextStart int, results [][]int) [][]int {
	if remain == 0 && len(comb) == k {
		cloned := make([]int, len(comb))
		copy(cloned, comb)
		results = append(results, cloned)
		return results
	} else if remain < 0 || len(comb) == k {
		return results
	}

	for i := nextStart; i < 9; i++ {
		comb = append(comb, i + 1)
		results = backtrack(remain - i - 1, k, comb, i + 1, results)
		comb = comb[:len(comb) - 1]
	}
	return results
}

func combinationSum3(k int, n int) [][]int {
	comb := []int{}
	results := [][]int{}

	results = backtrack(n, k, comb, 0, results)
	return results
}