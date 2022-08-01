package main

func uniquePaths(m int, n int) int {
	dp :=  make([][]int, m)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1
		}
	}
	for col := 1; col < m; col++ {
		for row := 1; row < n; row++ {
			dp[col][row] = dp[col - 1][row] + dp[col][row - 1]
		}
	}
	return dp[m - 1][n - 1]
}

