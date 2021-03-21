package main

func count(N int) []int {
	ans := make([]int, 10)
	for N > 0 {
		ans[N%10]++
		N /= 10
	}
	return ans
}

func equals(a []int, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func reorderedPowerOf2(N int) bool {
	A := count(N)
	for i := 0; i < 31; i++ {
		if equals(A, count(1<<i)) {
			return true
		}
	}
	return false
}
