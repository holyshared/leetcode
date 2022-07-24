package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])

	row := m - 1
	col := 0

	for row >= 0 && col < n {
		if matrix[row][col] > target {
			row--
		} else if matrix[row][col] < target {
			col++
		} else {
			return true
		}
	}

	return false
}
