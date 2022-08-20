package main

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func minRefuelStops(target int, startFuel int, stations [][]int) int {
	n := len(stations)
	dp := make([]int, n+1)
	dp[0] = startFuel

	for i := 0; i < n; i++ {
		for t := i; t >= 0; t-- {
			if dp[t] >= stations[i][0] {
				dp[t+1] = max(dp[t+1], dp[t]+stations[i][1])
			}
		}
	}

	for i := 0; i <= n; i++ {
		if dp[i] >= target {
			return i
		}
	}

	return -1
}
