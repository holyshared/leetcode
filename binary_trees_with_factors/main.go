package main

import (
	"fmt"
	"sort"
)

var mod = 1000000007

func numFactoredBinaryTrees(arr []int) int {
	sort.Slice(arr, func(i int, j int) bool {
		return arr[i] < arr[j]
	})

	num := len(arr)
	dp := make([]int, num)
	for i, _ := range dp {
		dp[i] = 1
	}

	index := map[int]int{}

	for i := 0; i < num; i++ {
		index[arr[i]] = i // value to index
	}

	/*
		        18
					6 * 3
				3 * 2

		        18
					6 * 3
				2 * 3

		        18
					3 *  6
				     3 * 2

		        18
					3 *  6
				     2 * 3

		        18
					9 * 2
				3 * 3

		        18
					2 * 9
				     3 * 3

	*/

	// [1, 1, 3, 7]
	// [2, 3, 6, 18]

	// i: 0, j: 0 to 0
	// i: 1, j: 0 to 1
	// i: 2, j: 0 to 2
	// i: 3, j: 0 to 3
	// i: 4, j: 0 to 4

	for i := 0; i < num; i++ {
		fmt.Printf("i=%d\n", i)
		for j := 0; j < i; j++ {
			fmt.Printf("i=%d, j=%d\n", i, j)
			fmt.Printf("%d, %d\n", arr[i], arr[j])
			if arr[i]%arr[j] == 0 {
				// example: 6 / 2 = 3
				right := arr[i] / arr[j]
				v, ok := index[right]
				if ok {
					dp[i] = (dp[i] + (dp[j] * dp[v])) % mod
				}
			}
		}
	}

	fmt.Println(dp)

	ans := 0
	for _, x := range dp {
		ans += x
	}

	return (int)(ans % mod)
}
