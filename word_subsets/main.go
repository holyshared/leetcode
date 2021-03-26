package main

import (
	"math"
)

func calcCharCount(word string) []int {
	ans := make([]int, 26)
	for _, c := range []rune(word) {
		ans[int(c-'a')]++
	}
	return ans
}

func wordSubsets(A []string, B []string) []string {
	bmax := calcCharCount("")

	for _, b := range B {
		bCount := calcCharCount(b)
		for i := 0; i < 26; i++ {
			bmax[i] = int(math.Max(float64(bmax[i]), float64(bCount[i])))
		}
	}

	ans := []string{}
	for _, a := range A {
		matchAll := true
		aCount := calcCharCount(a)
		for i := 0; i < 26; i++ {
			if aCount[i] < bmax[i] {
				matchAll = false
				break
			}
		}
		if matchAll {
			ans = append(ans, a)
		}
	}
	return ans
}
