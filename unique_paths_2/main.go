package main

/*
ロボットは、m x nグリッドの左上隅にあります（下の図で「開始」とマークされています）。
ロボットは、どの時点でも下または右にしか移動できません。
ロボットはグリッドの右下隅に到達しようとしています（下の図で「完了」とマークされています）。

次に、グリッドにいくつかの障害物が追加されているかどうかを検討します。 一意のパスはいくつありますか？

障害物とスペースは、グリッドでそれぞれ1と0としてマークされます。
*/

func _uniquePathsWithObstacles(obstacleGrid [][]int, x int, y int, ex int, ey int, curr int) int {
	if x == ex && y == ey {
		return curr + 1
	}
	if ey >= y+1 && obstacleGrid[x][y+1] == 0 {
		curr = _uniquePathsWithObstacles(obstacleGrid, x, y+1, ex, ey, curr)
	}
	if ex >= x+1 && obstacleGrid[x+1][y] == 0 {
		curr = _uniquePathsWithObstacles(obstacleGrid, x+1, y, ex, ey, curr)
	}
	return curr
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	maxX := len(obstacleGrid[0]) - 1
	maxY := len(obstacleGrid) - 1

	return _uniquePathsWithObstacles(obstacleGrid, 0, 0, maxX, maxY, 0)
}
