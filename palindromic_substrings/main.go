package main

import "strconv"

func isPalindromic(chars []rune) bool {
	if len(chars) == 1 {
		return true
	}
	left := 0
	right := len(chars) - 1

	for left < right {
		if chars[left] != chars[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func countSubstrings(s string) int {
	chars := []rune(s)
	l := len(chars)
	palindromic := map[string]bool{}

	count := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			prevKey := strconv.Itoa(i+1) + "-" + strconv.Itoa(j-1)
			palindromiced, ok := palindromic[prevKey]

			pass := false

			if ok {
				if palindromiced {
					if chars[i] == chars[j] {
						pass = true
					}
				}
			} else {
				pass = isPalindromic(chars[i:j])
			}

			if pass {
				count++
			}

			key := strconv.Itoa(i) + "-" + strconv.Itoa(j)
			palindromic[key] = true
		}
	}
	return count
}
