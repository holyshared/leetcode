package main

import (
	"math"
	"sort"
)

var mod = int(math.Pow10(9)) + 7

func threeSumMulti(arr []int, target int) int {
	ans := 0
	sort.Sort(sort.IntSlice(arr))
	l := len(arr)

	for i := 0; i < l; i++ {
		T := target - arr[i]
		j, k := i+1, l-1

		for j < k {
			if arr[j]+arr[k] < T {
				j++
			} else if arr[j]+arr[k] > T {
				k--
			} else if arr[j] != arr[k] {
				left, right := 1, 1
				for j+1 < k && arr[j] == arr[j+1] {
					left++
					j++
				}
				for k-1 > j && arr[k] == arr[k-1] {
					right++
					k--
				}

				ans += left * right
				ans %= mod
				j++
				k--
			} else {
				ans += (k - j + 1) * (k - j) / 2
				ans %= mod
				break
			}
		}
	}

	return ans
}
