/*
文字列のゴミのインデックスが 0 の配列が与えられます。

ここで、garbage[i] は i 番目の家のゴミの品揃えを表します。
ごみ[i]は、金属、紙、ガラスごみの1単位をそれぞれ表す文字「M」、「P」、「G」のみで構成されています。
どの種類のゴミでも 1 ユニット拾うのに 1 分かかります。

また、0 で始まる整数配列 travel も与えられます。ここで、travel[i] は家 i から家 i + 1 に移動するのに必要な分数です。

市内には 3 台のごみ収集車があり、それぞれが 1 種類のごみを収集します。
各ごみ収集車は家 0 から出発し、順番に各家を訪問する必要があります。 ただし、すべての家を訪問する必要はありません。

ごみ収集車は、一度に 1 台しか使用できません。 1台のトラックが運転中やゴミ拾いをしている間、他の2台のトラックは何もできません。

すべてのガベージをピックアップするのに必要な最小分数を返します。

garbage = ["G","P","GP","GG"], travel = [2,4,3] // 0->1, 1->2, 2->3
*/

package main

// 0, G:1
//      2
//    P:1
//      4
//      2
//      3
//      2
//
// g

var types = map[rune]int{
	'G': 0,
	'M': 1,
	'P': 2,
}

// types [
//   0 -> [0,2,3,]
//	 1 -> [0,2,3,]
//   2 -> [0,2,3,]
// ]

func garbageCollection(garbage []string, travel []int) int {
	remainsTypes := make([][][]int, 3)

	for i := 0; i < len(garbage); i++ {
		sums := make([]int, 3)
		for _, v := range garbage[i] {
			sums[types[v]]++
		}
		for _, ty := range types {
			if sums[ty] <= 0 {
				continue
			}
			if len(remainsTypes[ty]) <= 0 {
				remainsTypes[ty] = [][]int{{i, sums[ty]}}
			} else {
				remainsTypes[ty] = append(remainsTypes[ty], []int{i, sums[ty]})
			}
		}
	}

	ans := 0
	sums := make([]int, 3)
	for ty, remainGarbages := range remainsTypes {
		currHouse := 0
		if len(remainGarbages) <= 0 {
			continue
		}
		for _, countOfHouse := range remainGarbages {
			house, count := countOfHouse[0], countOfHouse[1]
			sums[ty] += count

			move := 0
			if currHouse != house {
				for j := currHouse; j < house; j++ {
					move += travel[j]
				}
			}
			sums[ty] += move
			currHouse = house
		}
		ans += sums[ty]
	}

	return ans
}
