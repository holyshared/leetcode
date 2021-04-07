package main

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
	'A': true,
	'E': true,
	'I': true,
	'O': true,
	'U': true,
}

func halvesAreAlike(s string) bool {
	l, r := 0, len(s)-1
	lv, rv := 0, 0
	chars := []rune(s)

	for l < r {
		_, lok := vowels[chars[l]]
		if lok {
			lv++
		}
		_, rok := vowels[chars[r]]
		if rok {
			rv++
		}
		l++
		r--
	}
	return lv == rv
}
