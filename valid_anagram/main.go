package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a, b := []rune(s), []rune(t)
	counts := make([]int, 26)

	for i := 0; i < len(a); i++ {
		ac, bc := a[i], b[i]
		counts[ac-'a']++
		counts[bc-'a']--
	}

	for _, count := range counts {
		if count != 0 {
			return false
		}
	}

	return true
}
