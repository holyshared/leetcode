package main

type CharCount map[rune]int

func calcCharCount(words []string) map[string]CharCount {
	charCount := map[string]CharCount{}
	for i := 0; i < len(words); i++ {
		chars := []rune(words[i])
		count := CharCount{}
		for _, c := range chars {
			curr, ok := count[c]
			if ok {
				count[c] = curr + 1
			} else {
				count[c] = 1
			}
		}
		charCount[words[i]] = count
	}
	return charCount
}

func isInclude(ac CharCount, bc CharCount) bool {
	for c, bcount := range bc {
		acount, ok := ac[c]
		// aにbがない
		if !ok {
			return false
		}
		if acount < bcount {
			return false
		}
	}
	return true
}

func wordSubsets(A []string, B []string) []string {
	a := calcCharCount(A)
	b := calcCharCount(B)

	results := []string{}

	for aw, awc := range a {
		matchAll := true
		for _, bwc := range b {
			if !isInclude(awc, bwc) {
				matchAll = false
				break
			}
		}
		if matchAll {
			results = append(results, aw)
		}
	}
	return results
}
