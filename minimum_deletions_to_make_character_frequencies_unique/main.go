package main

func countByChar(s string) map[rune]int {
	chars := []rune(s)
	counts := map[rune]int{}
	for i := 0; i < len(chars); i++ {
		counts[chars[i]] += 1
	}
	return counts
}

func minDeletions(s string) int {
	deleteCount := 0
	marked := map[int]bool{}

	countChars := countByChar(s)

	for _, count := range countChars {
		for count > 0 && marked[count] {
			count -= 1
			deleteCount += 1
		}
		marked[count] = true
	}
	return deleteCount
}
