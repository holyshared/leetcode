package main

func calculateTime(keyboard string, word string) int {
	keys := []rune(keyboard)
	wordChars := []rune(word)

	keyl := len(keys)
	keyMap := map[rune]int{}

	for i := 0; i < keyl; i++ {
		keyMap[keys[i]] = i
	}
	curr := 0
	step := 0

	for j := 0; j < len(wordChars); j++ {
		p, _ := keyMap[wordChars[j]]

		if curr < p {
			step += p - curr
		} else {
			step += curr - p
		}
		curr = p
	}
	return step
}
