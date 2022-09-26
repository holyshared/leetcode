package main

func reverse(s []rune, l, r int) []rune {
	for l < r {
		lc, rc := s[l], s[r]
		s[l] = rc
		s[r] = lc
		l++
		r--
	}
	return s
}

func reverseWords(s string) string {
	ans := []rune(s)
	left := 0
	for i, c := range ans {
		if c == ' ' {
			ans = reverse(ans, left, i-1)
			left = i + 1
		}
	}
	ans = reverse(ans, left, len(ans)-1)
	return string(ans)
}
