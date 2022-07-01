package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	total := 0
	remain := truckSize

	for i := 0; i < len(boxTypes); i++ {
		if boxTypes[i][0] < remain {
			remain -= boxTypes[i][0]
			total += (boxTypes[i][0] * boxTypes[i][1])
		} else {
			total += remain * boxTypes[i][1]
			break
		}
	}

	return total
}
