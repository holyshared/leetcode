package main

import (
	"math"
)

type ZeroCount int
type OneCount int

// 0,1の数を数える
func countZeroesOnes(s string) (ZeroCount, OneCount) {
	c := make([]int, 2)
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		c[int(runes[i]-'0')]++
	}
	return ZeroCount(c[0]), OneCount(c[1])
}

func calculate(strs []string, i int, zeroes int, ones int, memo [][][]int) int {
	if i == len(strs) {
		return 0
	}

	if memo[i][zeroes][ones] != 0 {
		return memo[i][zeroes][ones]
	}

	zeroCount, oneCount := countZeroesOnes(strs[i])
	taken := -1

	// 使用できる0,1の数が足りる場合
	remainZero := zeroes - int(zeroCount)
	remainOne := ones - int(oneCount)
	if remainZero >= 0 && remainOne >= 0 {
		// 1. 0,1が足りる場合はstrsの次の数値を調べる
		taken = calculate(strs, i+1, remainZero, remainOne, memo) + 1
	}
	// 2. 表現できない場合はスキップする
	notTaken := calculate(strs, i+1, zeroes, ones, memo)
	// 1か2で大きい方を採用する
	memo[i][zeroes][ones] = int(math.Max(float64(taken), float64(notTaken)))
	return memo[i][zeroes][ones]
}

/**
 * memo[strsの番号][0の数][1の数] = 最大値
 */
func findMaxForm(strs []string, m int, n int) int {
	memo := make([][][]int, len(strs))

	for i := 0; i < len(strs); i++ {
		mm := make([][]int, m+1)
		for j := 0; j < len(mm); j++ {
			nm := make([]int, n+1)
			mm[j] = nm
		}
		memo[i] = mm
	}
	return calculate(strs, 0, m, n, memo)
}
