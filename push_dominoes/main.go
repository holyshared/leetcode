package main

func pushDominoes(dominoes string) string {
	N := len(dominoes)
	chars := []rune(dominoes)
	indexes := make([]int, N+2)
	symbols := make([]rune, N+2)
	len := 1
	indexes[0] = -1
	symbols[0] = 'L'

	for i := 0; i < N; i++ {
		if chars[i] != '.' {
			indexes[len] = i
			symbols[len] = chars[i]
			len++
		}
	}

	indexes[len] = N
	symbols[len] = 'R'
	len++

	ans := []rune(dominoes)

	for index := 0; index < len-1; index++ {
		i, j := indexes[index], indexes[index+1]
		x, y := symbols[index], symbols[index+1]

		// L...L or R...R
		if x == y {
			for k := i + 1; k < j; k++ {
				ans[k] = x
			}
		} else if x > y { // R...L
			for k := i + 1; k < j; k++ {
				// R.L
				if k-i == j-k {
					ans[k] = '.'
				} else {
					if k-i < j-k {
						ans[k] = 'R'
					} else {
						ans[k] = 'L'
					}
				}
			}
		}
	}

	return string(ans)
}
