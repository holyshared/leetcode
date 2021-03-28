package main

import "strconv"

func countChars(s string) []int {
	counts := make([]int, 29)
	for _, c := range []rune(s) {
		counts[int(c-'a')] += 1
	}
	return counts
}

func originalDigits(s string) string {
	source := countChars(s)

	out := make([]int, 10)
	out[0] = source[int('z'-'a')] // zero
	out[2] = source[int('w'-'a')] // two
	out[4] = source[int('u'-'a')] // four
	out[6] = source[int('x'-'a')] // six
	out[8] = source[int('g'-'a')] // eight

	out[3] = source[int('h'-'a')] - out[8]                   // three
	out[5] = source[int('f'-'a')] - out[4]                   // five
	out[7] = source[int('s'-'a')] - out[6]                   // seven
	out[9] = source[int('i'-'a')] - out[5] - out[6] - out[8] // nine
	out[1] = source[int('n'-'a')] - out[7] - 2*out[9]

	result := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < out[i]; j++ {
			result += strconv.Itoa(i)
		}
	}
	return result
}
