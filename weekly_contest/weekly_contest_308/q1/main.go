package main

/*
長さ n の整数配列 nums と、長さ m の整数配列クエリが与えられます。

長さ m の配列 answer を返します。
ここで、answer[i] は、要素の合計が query[i] 以下になるように nums から取得できるサブシーケンスの最大サイズです。

サブシーケンスは、残りの要素の順序を変更せずに一部の要素を削除するか、要素をまったく削除しないことで、別の配列から派生できる配列です。


入力: 数値 = [4,5,2,1]、クエリ = [3,10,21]
出力: [2,3,4]
説明: 次のようにクエリに回答します。
- サブシーケンス [2,1] の合計は 3 以下です。このようなサブシーケンスの最大サイズは 2 であることが証明できるため、answer[0] = 2 です。
- サブシーケンス [4,5,1] の合計は 10 以下です。このようなサブシーケンスの最大サイズは 3 であることが証明できるため、answer[1] = 3 です。
- サブシーケンス [4,5,2,1] の合計は 21 以下です。このようなサブシーケンスの最大サイズは 4 であることが証明できるため、answer[2] = 4 です。
*/

import (
	"sort"
)

func find(q int, sums []int) int {
	res := sort.SearchInts(sums, q)
	if res > len(sums)-1 {
		return len(sums)
	} else {
		if q >= sums[res] {
			return res + 1
		} else {
			return res
		}
	}
}

func answerQueries(nums []int, queries []int) []int {
	ans := make([]int, len(queries))
	sums := make([]int, len(nums))
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sums[i] = sums[i-1] + nums[i]
	}

	for i, q := range queries {
		ans[i] = find(q, sums)
	}

	return ans
}
