package main

import "strings"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func alienOrder(words []string) string {
	adjList := map[rune][]rune{}
	counts := map[rune]int{}

	for _, word := range words {
		for _, c := range []rune(word) {
			counts[c] = 0
			adjList[c] = []rune{}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		if len(word1) > len(word2) && strings.HasPrefix(word1, word2) {
			return ""
		}

		for j := 0; j < min(len(word1), len(word2)); j++ {
			w1c, w2c := []rune(word1)[j], []rune(word2)[j]
			if w1c != w2c {
				adjList[w1c] = append(adjList[w1c], w2c)
				counts[w2c] = counts[w2c] + 1
				break
			}
		}
	}

	sb := []rune{}
	queue := []rune{}

	for k, _ := range counts {
		if counts[k] == 0 {
			queue = append(queue, k)
		}
	}

	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]

		sb = append(sb, c)

		for _, next := range adjList[c] {
			counts[next] = counts[next] - 1

			if counts[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(sb) < len(counts) {
		return ""
	}
	return string(sb)
}
