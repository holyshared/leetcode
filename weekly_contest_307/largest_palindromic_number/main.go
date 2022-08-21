package main

import (
	"strconv"
)

func largestPalindromic(num string) string {
	counts := make([]int, 10)
	for i := 0; i < len(num); i++ {
		counts[num[i]-'0']++
	}

	s := ""

	for i := 9; i >= 0; i-- {
		for j := 0; j < counts[i]/2; j++ {
			if i == 0 && len(s) == 0 {
				break
			}
			s += strconv.Itoa(i)
		}
	}

	rev := []rune(s)

	for i := 0; i < len(rev)/2; i++ {
		temp := rev[i]
		rev[i] = rev[len(rev)-i-1]
		rev[len(rev)-i-1] = temp
	}

	for i := 9; i >= 0; i-- {
		if counts[i]%2 == 1 {
			s += strconv.Itoa(i)
			break
		}
	}

	if len(s) == 0 {
		return "0"
	}
	return s + string(rev)
}
