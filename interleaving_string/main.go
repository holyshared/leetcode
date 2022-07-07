package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	a := []rune(s1)
	b := []rune(s2)
	c := []rune(s3)

	if len(a)+len(b) != len(c) {
		return false
	}

	dp := make([][]bool, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]bool, len(b)+1)
	}

	for i := 0; i <= len(a); i++ {
		for j := 0; j <= len(b); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] && b[j-1] == c[i+j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] && a[i-1] == c[i+j-1]
			} else {
				dp[i][j] = (dp[i-1][j] && a[i-1] == c[i+j-1]) || (dp[i][j-1] && b[j-1] == c[i+j-1])
			}
		}
	}
	return dp[len(a)][len(b)]
}
