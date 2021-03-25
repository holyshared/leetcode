package main

var DIRECTIONS = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

type Solution struct {
	numRows     int
	numCols     int
	landHeights [][]int
}

func NewSolution(landHeights [][]int) *Solution {
	return &Solution{
		numRows:     len(landHeights),
		numCols:     len(landHeights[0]),
		landHeights: landHeights,
	}
}

func (this *Solution) PacificAtlantic() [][]int {
	matrix := this.landHeights
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	this.numRows = len(matrix)
	this.numCols = len(matrix[0])
	this.landHeights = matrix

	pacificQueue := [][]int{}
	atlanticQueue := [][]int{}

	for i := 0; i < this.numRows; i++ {
		pacificQueue = append(pacificQueue, []int{i, 0})
		atlanticQueue = append(atlanticQueue, []int{i, this.numCols - 1})
	}

	for i := 0; i < this.numCols; i++ {
		pacificQueue = append(pacificQueue, []int{0, i})
		atlanticQueue = append(atlanticQueue, []int{this.numRows - 1, i})
	}

	pacificReachable := this.bfs(pacificQueue)
	atlanticReachable := this.bfs(atlanticQueue)

	commonCells := [][]int{}

	for i := 0; i < this.numRows; i++ {
		for j := 0; j < this.numCols; j++ {
			if pacificReachable[i][j] && atlanticReachable[i][j] {
				commonCells = append(commonCells, []int{i, j})
			}
		}
	}
	return commonCells
}

func (this *Solution) bfs(queue [][]int) [][]bool {
	reachable := make([][]bool, this.numRows)
	for i := 0; i < this.numRows; i++ {
		reachable[i] = make([]bool, this.numCols)
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		reachable[cell[0]][cell[1]] = true

		for _, dir := range DIRECTIONS {
			newRow := cell[0] + dir[0]
			newCol := cell[1] + dir[1]

			if newRow < 0 || newRow >= this.numRows || newCol < 0 || newCol >= this.numCols {
				continue
			}
			if reachable[newRow][newCol] {
				continue
			}
			// 新しいセルの高さが同じかそれ以上であることを確認します。
			// 水が新しいセルから古いセルに流れることができるように
			if this.landHeights[newRow][newCol] < this.landHeights[cell[0]][cell[1]] {
				continue
			}
			queue = append(queue, []int{newRow, newCol})
		}
	}
	return reachable
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	solution := NewSolution(matrix)
	return solution.PacificAtlantic()
}
