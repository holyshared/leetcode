package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var digits = map[rune]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

func isDigit(c rune) bool {
	_, ok := digits[c]
	return ok
}

func getNumberAt(s []rune, index int) (int, int) {
	negative := false

	if s[index] == '-' {
		negative = true
		index++
	}

	number := 0
	numberRunes := []rune{}
	for index < len(s) && isDigit(s[index]) {
		numberRunes = append(numberRunes, s[index])
		index++
	}
	iv, _ := strconv.Atoi(string(numberRunes))
	number = iv

	if negative {
		return -number, index
	} else {
		return number, index
	}
}

func str2tree(s string) *TreeNode {
	if s == "" {
		return nil
	}

	chars := []rune(s)

	root := &TreeNode{Val: -1}
	stack := []*TreeNode{root}

	index := 0
	for index < len(chars) {
		node := stack[len(stack)-1]
		stack = stack[:(len(stack) - 1)]

		if isDigit(chars[index]) || chars[index] == '-' {
			num, lastIndex := getNumberAt(chars, index)
			index = lastIndex

			node.Val = num

			if index < len(chars) && chars[index] == '(' {
				stack = append(stack, node)
				node.Left = &TreeNode{Val: -1}
				stack = append(stack, node.Left)
			}
		} else if chars[index] == '(' && node.Left != nil {
			stack = append(stack, node)
			node.Right = &TreeNode{Val: -1}
			stack = append(stack, node.Right)
		}
		index++
	}
	if len(stack) <= 0 {
		return root
	} else {
		return stack[len(stack)-1]
	}
}
