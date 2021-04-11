package main

// 0            1
// 1          1   1
// 2        1   2   1
// 3      1   3   3   1
// 4    1   4   6   4  1
// 5  1   5   10  10   5  1

func _getRow(maxLevel int, level int, levels [][]int) [][]int {
	if level >= maxLevel {
		return levels
	}
	a := make([]int, level+1)
	levels[level] = a

	levels[level][0] = 1
	levels[level][len(a)-1] = 1

	j := 0
	for i := 1; i < len(a)-1; i++ {
		levels[level][i] = levels[level-1][j] + levels[level-1][j+1]
		j++
	}
	return _getRow(maxLevel, level+1, levels)
}

func getRow(rowIndex int) []int {
	levels := make([][]int, rowIndex+1)
	levels[0] = []int{1}
	levels = _getRow(rowIndex+1, 1, levels)

	return levels[rowIndex]
}
