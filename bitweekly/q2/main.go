package main

import "strconv"

/*
正の整数 n を指定すると、合計が n になる 2 のべき乗の最小数で構成される、powers と呼ばれる 0 から始まる配列が存在します。
配列は非降順でソートされ、配列を形成する方法は 1 つしかありません。

また、query[i] = [lefti, righti] である 0 インデックスの 2D 整数配列クエリも与えられます。
各 query[i] は、lefti <= j <= righti ですべての powers[j] の積を見つけなければならないクエリを表します。

配列の回答を返します。長さはクエリと同じです。
ここで、answers[i] は i 番目のクエリに対する回答です。
i 番目のクエリに対する回答が大きすぎる可能性があるため、各回答 [i] は 109 + 7 を法として返される必要があります。

Input: n = 15, queries = [[0,1],[2,2],[0,3]]
Output: [2,4,64]
Explanation:
For n = 15, powers = [1,2,4,8]. It can be shown that powers cannot be a smaller size.
Answer to 1st query: powers[0] * powers[1] = 1 * 2 = 2.
Answer to 2nd query: powers[2] = 4.
Answer to 3rd query: powers[0] * powers[1] * powers[2] * powers[3] = 1 * 2 * 4 * 8 = 64.
Each answer modulo 109 + 7 yields the same answer, so [2,4,64] is returned.

*/

func productQueries(n int, queries [][]int) []int {
	curr := 0
	powers := []int{}
	memo := map[string]int{}

	for i := 0; n > curr; i++ {
		if n&(1<<i) != 0 {
			curr += 1 << i
			powers = append(powers, 1<<i)
		}
	}

	answers := make([]int, len(queries))
	for i, q := range queries {
		f, t := q[0], q[1]
		key := strconv.Itoa(f) + ":" + strconv.Itoa(t)

		val, has := memo[key]
		if has {
			answers[i] = val
		} else {
			res := 1
			for si := f; si <= t; si++ {
				res = (res * powers[si]) % 1000000007
			}
			answers[i] = res
			memo[key] = res
		}
	}
	return answers
}
