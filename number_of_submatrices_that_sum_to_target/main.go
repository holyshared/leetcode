package main

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	r, c := len(matrix), len(matrix[0])

	ps := make([][]int, r+1)
	for i := 0; i < len(ps); i++ {
		ps[i] = make([]int, c+1)
	}

	for i := 1; i < r+1; i++ {
		for j := 1; j < c+1; j++ {
			ps[i][j] = ps[i-1][j] + ps[i][j-1] - ps[i-1][j-1] + matrix[i-1][j-1]
		}
	}

	count, currSum := 0, 0
	h := map[int]int{}

	for r1 := 1; r1 < r+1; r1++ {
		for r2 := r1; r2 < r+1; r2++ {
			h = map[int]int{}
			h[0] = 1

			for col := 1; col < c+1; col++ {
				currSum = ps[r2][col] - ps[r1-1][col]

				v, ok := h[currSum-target]
				if ok {
					count += v
				}

				curr, hasCurr := h[currSum]
				if hasCurr {
					h[currSum] = curr + 1
				} else {
					h[currSum] += 1
				}
			}
		}
	}

	return count
}
