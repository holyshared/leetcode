package main

func runeOrders(order string) map[rune]int {
	runeOrders := map[rune]int{}
	scoreOfRunes := []rune(order)

	for i, c := range scoreOfRunes {
		runeOrders[c] = i
	}
	return runeOrders
}

func isAlienSorted(words []string, order string) bool {
	sortedWords := make([]string, len(words))
	copy(sortedWords, words)

	orders := runeOrders(order)

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			arunes := []rune(words[i])
			brunes := []rune(words[i+1])

			if j >= len(words[i+1]) {
				return false
			}

			if arunes[j] != brunes[j] {
				as, _ := orders[arunes[j]]
				bs, _ := orders[brunes[j]]

				if as > bs {
					return false
				} else {
					break
				}
			}
		}
	}

	return true
}
