package main

func longestValidParentheses(s string) int {
	if len(s) <= 0 {
		return 0
	}
	stack := []int{}

	// (()
	chars := []rune(s)
	max := 0
	stack = append(stack, -1)
	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:(len(stack) - 1)]
			if len(stack) <= 0 {
				stack = append(stack, i)
			} else {
				if max < i-stack[len(stack)-1] {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}

	return max
}
