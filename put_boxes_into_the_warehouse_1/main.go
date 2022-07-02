package main

import (
	"sort"
)

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	// 一つ前のより、後の方が大きい場合。一つ前に合わせる
	// 後が大きくても、手前が小さいので手前の大きさと同等のものしか箱を押し込めない
	for i := 1; i < len(warehouse); i++ {
		if warehouse[i-1] < warehouse[i] {
			warehouse[i] = warehouse[i-1]
		}
	}

	sort.Slice(boxes, func(i, j int) bool {
		return boxes[i] < boxes[j]
	})

	count := 0

	for i := len(warehouse) - 1; i >= 0; i-- {
		if count < len(boxes) && boxes[count] <= warehouse[i] {
			count += 1
		}
	}

	return count
}
