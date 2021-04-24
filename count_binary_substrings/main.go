package main

func countBinarySubstrings(s string) int {
	ans, prev, cur := 0, 0, 1
	runes := []rune(s)

	for i := 1; i < len(s); i++ {
		if runes[i-1] != runes[i] {

			if prev > cur {
				ans += cur
			} else {
				ans += prev
			}
			prev = cur
			cur = 1
		} else {
			cur++
		}
	}
	if prev > cur {
		ans += cur
	} else {
		ans += prev
	}
	return ans
}
