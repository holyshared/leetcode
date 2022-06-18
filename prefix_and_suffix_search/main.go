package main

import (
	"strings"
)

type WordFilter struct {
	words []string
}

func Constructor(words []string) WordFilter {
	return WordFilter{
		words: words,
	}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	lastWordIndex := len(this.words) - 1

	for i := lastWordIndex; i >= 0; i-- {
		matched := strings.HasPrefix(this.words[i], prefix) && strings.HasSuffix(this.words[i], suffix)
		if matched {
			return i
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
