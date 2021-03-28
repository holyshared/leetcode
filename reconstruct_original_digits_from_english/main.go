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
	out[0] = source[int('z'-'a')]
	out[2] = source[int('w'-'a')]
	out[4] = source[int('u'-'a')]
	out[6] = source[int('x'-'a')]
	out[8] = source[int('g'-'a')]

	out[3] = source[int('h'-'a')] - out[8]
	out[5] = source[int('f'-'a')] - out[4]
	out[7] = source[int('s'-'a')] - out[6]
	out[9] = source[int('i'-'a')] - out[5] - out[6] - out[8]
	out[1] = source[int('n'-'a')] - out[7] - 2*out[9]

	result := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < out[i]; j++ {
			result += strconv.Itoa(i)
		}
	}
	return result
}
