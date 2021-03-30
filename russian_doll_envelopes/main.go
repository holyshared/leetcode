package main

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		} else {
			return envelopes[i][0] < envelopes[j][0]
		}
	})

	dp := []int{}
	for _, envelope := range envelopes {
		// SearchIntsは、ソートされたintのスライスでxを検索し、Searchで指定されたインデックスを返します。
		// 戻り値は、xが存在しない場合にxを挿入するためのインデックスです（len（a）の可能性があります）。
		// スライスは昇順で並べ替える必要があります。
		// https://golang.org/pkg/sort/#SearchInts
		i := sort.SearchInts(dp, envelope[1])

		if i == len(dp) {
			// 一番大きいので最後に追加する
			dp = append(dp, envelope[1])
		} else {
			dp[i] = envelope[1]
		}
	}
	return len(dp)
}
