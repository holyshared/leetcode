package main

func reversed(s string) string {
	n := len(s)
	runes := []rune(s)
	for i := 0; i < n/2; i++ {
		a, b := runes[i], runes[n-i-1]
		runes[i] = b
		runes[n-i-1] = a
	}
	return string(runes)
}

func isPalindromeBetween(word string, front int, back int) bool {
	for front < back {
		if word[front] != word[back] {
			return false
		}
		front++
		back--
	}
	return true
}

func allValidPrefixes(word string) []string {
	validPrefixes := []string{}
	for i := 0; i < len(word); i++ {
		if isPalindromeBetween(word, i, len(word)-1) {
			validPrefixes = append(validPrefixes, string(word[0:i]))
		}
	}
	return validPrefixes
}

func allValidSuffixes(word string) []string {
	validSuffixes := []string{}
	for i := 0; i < len(word); i++ {
		if isPalindromeBetween(word, 0, i) {
			validSuffixes = append(validSuffixes, string(word[(i+1):]))
		}
	}
	return validSuffixes
}

func palindromePairs(words []string) [][]int {
	wordSet := map[string]int{}
	for i := 0; i < len(words); i++ {
		wordSet[words[i]] = i
	}
	solution := [][]int{}

	for _, word := range words {
		currentWordIndex := wordSet[word]
		reversedWord := reversed(word)

		val, has := wordSet[reversedWord]
		if has && val != currentWordIndex {
			solution = append(solution, []int{currentWordIndex, val})
		}

		for _, suffix := range allValidSuffixes(word) {
			reversedSuffix := reversed(suffix)
			val, has := wordSet[reversedSuffix]
			if has {
				solution = append(solution, []int{val, currentWordIndex})
			}
		}

		for _, prefix := range allValidPrefixes(word) {
			reversedPrefix := reversed(prefix)
			val, has := wordSet[reversedPrefix]
			if has {
				solution = append(solution, []int{currentWordIndex, val})
			}
		}
	}

	return solution
}
