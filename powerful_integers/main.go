package main

import (
	"math"
	"sort"
)

func powerfulIntegers(x int, y int, bound int) []int {
	a, b := 0, 0

	if x == 1 {
		a = bound
	} else {
		a = int(math.Log(float64(bound)) / math.Log(float64(x)))
	}
	if y == 1 {
		b = bound
	} else {
		b = int(math.Log(float64(bound)) / math.Log(float64(y)))
	}

	powerfulIntegers := map[int]bool{}

	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			value := int(math.Pow(float64(x), float64(i))) + int(math.Pow(float64(y), float64(j)))

			if value <= bound {
				powerfulIntegers[value] = true
			}

			if y == 1 {
				break
			}
		}

		if x == 1 {
			break
		}
	}

	results := []int{}
	for v, _ := range powerfulIntegers {
		results = append(results, v)
	}
	sort.Sort(sort.IntSlice(results))

	return results
}
