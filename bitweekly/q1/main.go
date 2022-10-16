package main

import "strconv"

/*
time と呼ばれる長さ 5 の文字列が与えられます。
これは、デジタル時計の現在の時刻を "hh:mm" の形式で表します。
最も早い時刻は「00:00」、最も遅い時刻は「23:59」です。

文字列時間では、? で表される数字は 記号は不明であり、0 から 9 までの数字に置き換える必要があります。

? ごとに置き換えることで作成できる有効なクロック時間の数である整数の回答を返します 0 から 9 までの数字。

?5:00

*/

func countTime(time string) int {
	hours1 := rune(time[0])
	hours2 := rune(time[1])
	mitunes1 := rune(time[3])
	mitunes2 := rune(time[4])

	ans := 1

	if hours1 == '?' && hours2 == '?' {
		ans = ans * 24
	} else if hours1 == '?' && hours2 != '?' {
		x, _ := strconv.Atoi(string(hours2))
		if x >= 4 {
			ans = ans * 2
		} else {
			ans = ans * 3
		}
	} else if hours1 != '?' && hours2 == '?' {
		x, _ := strconv.Atoi(string(hours1))
		if x == 2 {
			ans = ans * 4
		} else {
			ans = ans * 10
		}
	}

	// "?2:16"
	// 012345
	if mitunes1 == '?' && mitunes2 == '?' {
		ans = ans * 60
	} else if mitunes1 == '?' && mitunes2 != '?' {
		ans = ans * 6
	} else if mitunes1 != '?' && mitunes2 == '?' {
		ans = ans * 10
	}

	return ans
}
