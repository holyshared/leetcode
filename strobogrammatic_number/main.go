package main

func isStrobogrammatic(num string) bool {
	numRunes := []rune(num)
	revNumRunes := []rune{}

	for i := len(numRunes) - 1; i >= 0; i-- {
		c := numRunes[i]
		if c == '1' || c == '8' || c == '0' {
			revNumRunes = append(revNumRunes, c)
		} else if c == '6' {
			revNumRunes = append(revNumRunes, '9')
		} else if c == '9' {
			revNumRunes = append(revNumRunes, '6')
		} else {
			return false
		}
	}
	return string(revNumRunes) == num
}
