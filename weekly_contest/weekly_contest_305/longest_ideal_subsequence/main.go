package main

import (
	"math"
)

func maxOf(values ...int) int {
	max := 0
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}

func longestIdealString(s string, k int) int {
	//	ans := 1
	dp := make([]int, 26)

	n := len(s)
	for i := 0; i < n; i++ {
		v := int(s[i] - 'a')
		t := dp[v]
		for c := 0; c < 26; c++ {
			if int(math.Abs(float64(c-v))) <= k {
				t = maxOf(t, dp[c]+1)
			}
		}
		dp[v] = t
	}

	return maxOf(dp...)
}
