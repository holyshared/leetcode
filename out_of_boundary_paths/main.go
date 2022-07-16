package main

var mod int64 = 1000000007

type Finder struct {
	m int
	n int
	memo [][][]int64
}

func (this *Finder) dfs(i, j, remain int) int64 {
	// 範囲外に出た
	if i < 0 || j < 0 || this.m <= i || this.n <= j {
		return 1
		return 0
	} else if remain <= 0 {
		return 0
	}

	if this.memo[i][j][remain] >= 0 {
		return this.memo[i][j][remain]
	}

	v1 := this.dfs(i - 1, j, remain - 1)
	v2 := this.dfs(i + 1, j, remain - 1)
	v3 := this.dfs(i, j + 1, remain - 1)
	v4 := this.dfs(i, j - 1, remain - 1)

	this.memo[i][j][remain] = (((v1 + v2) % mod) + ((v3 + v4) % mod)) % mod

	return this.memo[i][j][remain]
}


func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	memo := make([][][]int64, m)
	for i := 0; i < m; i++ {
		memo[i] = make([][]int64, n)
		for j := 0; j < n; j++ {
			memo[i][j] = make([]int64, maxMove + 1)
			for k := 0; k <= maxMove; k++ {
				memo[i][j][k] = -1
			}
		}
	}

	finder := &Finder{m, n, memo}
	result := finder.dfs(startRow, startColumn, maxMove)

	return (int)(result)
}