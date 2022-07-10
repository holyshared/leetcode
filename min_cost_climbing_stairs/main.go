package main


func calculate(i int, cost []int, memo map[int]int) int {
	top := len(cost) - 1

	if i + 1 > top || i + 2 > top {
		return 0
	}

	value, ok := memo[i]
	if ok {
		return value
	}

	a := cost[i + 1] + calculate(i + 1, cost, memo)
	b := cost[i + 2] + calculate(i + 2, cost, memo)

	if a > b {
		memo[i] = b
		return b
	} else {
		memo[i] = a
		return a
	}
}

func minCostClimbingStairs(cost []int) int {
	memo := map[int]int{}
	a := cost[0] + calculate(0, cost, memo)
	b := cost[1] + calculate(1, cost, memo)

	if a > b {
		return b
	} else {
		return a
	}
}