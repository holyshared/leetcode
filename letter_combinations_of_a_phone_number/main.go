package main

var numberRunes = map[rune][]rune{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func _letterCombinations(numbers []rune, curr int, comb []rune, out []string) []string {
	if curr >= len(numbers) {
		return append(out, string(comb))
	}
	num := numbers[curr]
	runes, _ := numberRunes[num]

	for _, c := range runes {
		comb = append(comb, c)
		out = _letterCombinations(numbers, curr+1, comb, out)
		comb = comb[:(len(comb) - 1)]
	}
	return out
}

func letterCombinations(digits string) []string {
	if len(digits) <= 0 {
		return []string{}
	}

	chars := []rune(digits)

	num := chars[0]
	startRunes, _ := numberRunes[num]
	results := []string{}

	for _, c := range startRunes {
		comb := []rune{c}
		results = _letterCombinations(chars, 1, comb, results)
	}
	return results
}
