package main

/*
	card points: [1, 2, 3, 4, 5, 6, 1]
	選択枚数: 3

	前: [0, 1, 3, 6]  前から順番に選択した枚数の合計
	後: [0, 1, 7, 12] 後から順番に選択した枚数の合計

	下記の順番で最大となるスコアを決める
	1. 後ろから3つ選択した場合のスコア
	2. 前から1、後ろから2つ選択した場合のスコア
  3. 前から2、後ろから1つ選択した場合のスコア
  4. 前から3つ選択した場合のスコア
*/

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	frontSetOfCards := make([]int, k+1)
	rearSetOfCards := make([]int, k+1)

	for i := 0; i < k; i++ {
		frontSetOfCards[i+1] = cardPoints[i] + frontSetOfCards[i]
		rearSetOfCards[i+1] = cardPoints[n-i-1] + rearSetOfCards[i]
	}

	maxScore := 0

	for i := 0; i <= k; i++ {
		currentScore := frontSetOfCards[i] + rearSetOfCards[k-i]

		if maxScore < currentScore {
			maxScore = currentScore
		}
	}

	return maxScore
}
