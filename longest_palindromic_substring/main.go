package main

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 0 {
		return ""
	}

	originalRunes := []rune(s)
	reversedRunes := make([]rune, len(originalRunes))

	j := length - 1
	for i := 0; i < length; i++ {
		reversedRunes[i] = originalRunes[j]
		j--
	}

	dp := make([][]int, len(s))
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}

	maxLen := 0
	maxEnd := 0

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if originalRunes[i] == reversedRunes[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}

			if dp[i][j] > maxLen {
				if i-dp[i][j]+1 == length-1-j {
					maxLen = dp[i][j]
					maxEnd = i
				}
			}
		}
	}

	lcs := originalRunes[maxEnd-maxLen+1 : maxEnd+1]

	return string(lcs)
}
