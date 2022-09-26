package main

var mod = 1000000007

func concatenatedBinary(n int) int {
	ans := 0
	length := 0

	for i := 1; i <= n; i++ {
		if (i & (i - 1)) == 0 {
			length++
		}
		ans = ((ans << length) | i) % mod
	}
	return ans
}
