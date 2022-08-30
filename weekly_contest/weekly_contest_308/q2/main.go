/*
星印 * を含む文字列 s が与えられます。

1 回の操作で、次のことができます。

s の星を選択してください。
その左に最も近い星以外の文字を削除し、星自体も削除します。
すべての星が削除された後に文字列を返します。

ノート：

入力は、操作が常に可能になるように生成されます。
結果の文字列が常に一意であることを示すことができます。
*/

/*
Input: s = "leet**cod*e"

leet**cod*e

Output: "lecoe"
*/

package main

func removeStars(s string) string {
	chars := []rune(s)
	ans := []rune{}

	for i := 0; i < len(chars); i++ {
		if chars[i] != '*' {
			ans = append(ans, chars[i])
			continue
		} else {
			ans = ans[:len(ans)-1]
		}
	}

	return string(ans)
}
