package main

import "sort"

/*
文字列名の配列と、異なる正の整数で構成される高さの配列が与えられます。 両方の配列の長さは n です。

各インデックス i について、names[i] と heights[i] は i 番目の人物の名前と身長を示します。

身長の高い順に名前を並べ替えて返します。
*/

func sortPeople(names []string, heights []int) []string {
	values := make([][]int, len(heights))
	for i, height := range heights {
		values[i] = []int{i, height}
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i][1] > values[j][1]
	})
	ans := []string{}
	for _, pair := range values {
		i := pair[0]
		ans = append(ans, names[i])
	}
	return ans
}
