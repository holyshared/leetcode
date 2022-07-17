package main

var mod = 1000000007

type Inverse struct {
	memo [][]int
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func (this *Inverse) kInversePairs(n int, k int) int {
	if n == 0 {
		return 0
	}

	if k == 0 {
		return 1
	}

	if this.memo[n][k] != -1 {
		return this.memo[n][k]
	}

	inv := 0
	for i := 0; i <= min(k, n-1); i++ {
		inv = (inv + this.kInversePairs(n-1, k-i)) % mod
	}
	this.memo[n][k] = inv
	return inv
}

func kInversePairs(n int, k int) int {
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, k+1)
		for j := 0; j < k+1; j++ {
			memo[i][j] = -1
		}
	}

	finder := &Inverse{memo}
	return finder.kInversePairs(n, k)
}
