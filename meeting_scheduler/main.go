package main

import "sort"

/*

2人の可用性タイムスロット配列slot1とslots2、および会議の継続時間が与えられた場合、
両方で機能し、継続時間の最も早いタイムスロットを返します。

要件を満たす共通のタイムスロットがない場合は、空の配列を返します。

タイムスロットの形式は、開始から終了までの包括的な時間範囲を表す2つの要素[開始、終了]の配列です。

同じ人の2つの可用性スロットが互いに交差しないことが保証されています。
つまり、同じ人物の任意の2つのタイムスロット[start1、end1]と[start2、end2]について、start1> end2またはstart2> end1のいずれかです。

*/

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.SliceStable(slots1, func(i, j int) bool {
		return slots1[i][0] < slots1[j][0]
	})
	sort.SliceStable(slots2, func(i, j int) bool {
		return slots2[i][0] < slots2[j][0]
	})

	pointer1, pointer2 := 0, 0
	intersectLeft, intersectRight := 0, 0

	for pointer1 < len(slots1) && pointer2 < len(slots2) {
		if slots1[pointer1][0] < slots2[pointer2][0] {
			intersectLeft = slots2[pointer2][0]
		} else {
			intersectLeft = slots1[pointer1][0]
		}

		if slots1[pointer1][1] > slots2[pointer2][1] {
			intersectRight = slots2[pointer2][1]
		} else {
			intersectRight = slots1[pointer1][1]
		}

		if intersectRight-intersectLeft >= duration {
			return []int{intersectLeft, intersectLeft + duration}
		}

		if slots1[pointer1][1] < slots2[pointer2][1] {
			pointer1++
		} else {
			pointer2++
		}
	}
	return []int{}
}
