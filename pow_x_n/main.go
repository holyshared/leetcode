package main

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}

func myPow(x float64, n int) float64 {
	N := n
	if N < 0 {
		x = 1 / x
		N = -N
	}
	return fastPow(x, N)
}
