package main

func findAndReplacePattern(words []string, pattern string) []string {
	ans := []string{}
	for _, word := range words {
		if match(word, pattern) {
			ans = append(ans, word)
		}
	}
	return ans
}

func match(word, pattern string) bool {
	m1, m2 := map[rune]rune{}, map[rune]rune{}

	pt := []rune(pattern)

	for i, w := range []rune(word) {
		p := pt[i]

		if _, ok := m1[w]; !ok {
			m1[w] = p
		}

		if _, ok := m2[p]; !ok {
			m2[p] = w
		}
		m1w := m1[w]
		m2p := m2[p]
		if m1w != p || m2p != w {
			return false
		}
	}
	return true
}
