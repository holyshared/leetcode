package main

/*
文字列 s と、現在空の文字列 t を保持しているロボットが与えられます。 s と t が両方とも空になるまで、次の操作のいずれかを適用します。

文字列 s の最初の文字を削除し、ロボットに渡します。 ロボットは、この文字を文字列 t に追加します。
文字列 t の最後の文字を削除し、ロボットに渡します。 ロボットはこの文字を紙に書きます。
紙に書くことができる辞書編集的に最小の文字列を返します。


Input: s = "zza"
Output: "azz"
Explanation: Let p denote the written string.
Initially p="", s="zza", t="".
Perform first operation three times p="", s="", t="zza".
Perform second operation three times p="azz", s="", t="".


*/

func reverse(chars []rune) string {
	s := len(chars)
	for i := 0; i < s/2; i++ {
		l, r := chars[i], chars[s-i-1]
		chars[i] = r
		chars[s-i-1] = l
	}
	return string(chars)
}

func robotWithString(s string) string {
	chars := []rune(s)
	min := 27
	mini := 0
	for i := 0; i < len(chars); i++ {
		a := int(chars[i] - 'a')
		if min > a {
			min = a
			mini = i
		}
	}
	t := chars[0 : mini+1]
	r := chars[mini+1:]

	return reverse(t) + reverse(r)
}
