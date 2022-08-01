package main

func findPermutation(s string) []int {
	res := make([]int, len(s)+1)
	stack := []int{}

	j := 0
	runes := []rune(s)
	for i := 1; i <= len(s); i++ {
		if runes[i-1] == 'I' {
			stack = append(stack, i)
			for len(stack) > 0 {
				last := len(stack) - 1
				curr := stack[last]
				stack = stack[:last]
				res[j] = curr
				j++
			}
		} else {
			stack = append(stack, i)
		}
	}

	stack = append(stack, len(s)+1)

	for len(stack) > 0 {
		last := len(stack) - 1
		curr := stack[last]
		stack = stack[:last]
		res[j] = curr
		j++
	}

	return res
}
