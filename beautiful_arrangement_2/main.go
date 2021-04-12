package main

func constructArray(n int, k int) []int {
	c := 0
	ans := make([]int, n)
	for v := 1; v < n-k; v++ {
		ans[c] = v
		c++
	}
	for i := 0; i <= k; i++ {
		if i%2 == 0 {
			ans[c] = (n - k + i/2)
		} else {
			ans[c] = (n - i/2)
		}
		c++
	}
	return ans
}
