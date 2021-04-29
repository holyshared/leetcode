package main

/*
ロボットは、m x nグリッドの左上隅にあります（下の図で「開始」とマークされています）。
ロボットは、どの時点でも下または右にしか移動できません。
ロボットはグリッドの右下隅に到達しようとしています（下の図で「完了」とマークされています）。

次に、グリッドにいくつかの障害物が追加されているかどうかを検討します。 一意のパスはいくつありますか？

障害物とスペースは、グリッドでそれぞれ1と0としてマークされます。
*/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	R := len(obstacleGrid)
	C := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	obstacleGrid[0][0] = 1

	for i := 1; i < R; i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 1 {
			obstacleGrid[i][0] = 1
		} else {
			obstacleGrid[i][0] = 0
		}
	}

	for i := 1; i < C; i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 1 {
			obstacleGrid[0][i] = 1
		} else {
			obstacleGrid[0][i] = 0
		}
	}

	for i := 1; i < R; i++ {
		for j := 1; j < C; j++ {
			if obstacleGrid[i][j] == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			} else {
				obstacleGrid[i][j] = 0
			}
		}
	}

	return obstacleGrid[R-1][C-1]
}
