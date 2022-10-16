package main

import (
	"math"
)

/*
n 人の従業員がいて、それぞれが 0 から n - 1 までの一意の ID を持っています。

log[i] = [idi, leaveTimei] ここで、2D 整数配列ログが与えられます。

idi は、i 番目のタスクに取り組んだ従業員の ID です。
leaveTimei は、従業員が i 番目のタスクを終了した時刻です。 すべての値 leaveTimei は一意です。
i 番目のタスクは (i - 1) 番目のタスクが終了した直後に開始され、0 番目のタスクは時刻 0 に開始されることに注意してください。

タスクに最も長く取り組んだ従業員の ID を返します。 2 人以上の従業員が同数の場合、その中で最小の ID を返します。
*/

func hardestWorker(n int, logs [][]int) int {
	longTimeId := n
	maxProessTime := math.MinInt32
	startAt := 0

	for _, time := range logs {
		id, lt := time[0], time[1]
		tt := lt - startAt
		if maxProessTime <= tt {
			if maxProessTime == tt {
				if longTimeId > id {
					longTimeId = id
				}
			} else {
				longTimeId = id
				maxProessTime = tt
			}
		}
		startAt = lt
	}

	return longTimeId
}
