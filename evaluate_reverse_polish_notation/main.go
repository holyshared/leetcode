package main

import (
	"strconv"
)

type Solution struct {
	pos    int
	tokens []string
}

func (this Solution) computeOp() []string {
	op := this.tokens[this.pos]
	left, _ := strconv.Atoi(this.tokens[this.pos-2])
	right, _ := strconv.Atoi(this.tokens[this.pos-1])

	if op == "+" {
		this.tokens[this.pos-2] = strconv.Itoa(left + right)
	} else if op == "-" {
		this.tokens[this.pos-2] = strconv.Itoa(left - right)
	} else if op == "/" {
		this.tokens[this.pos-2] = strconv.Itoa(left / right)
	} else {
		this.tokens[this.pos-2] = strconv.Itoa(left * right)
	}

	return append(this.tokens[0:(this.pos-1)], this.tokens[(this.pos+1):]...)
}

func (this Solution) eval() int {
	for len(this.tokens) > 1 {
		token := this.tokens[this.pos]
		if token == "+" || token == "-" || token == "/" || token == "*" {
			this.tokens = this.computeOp()
			this.pos -= 1
		} else {
			this.pos++
		}
	}
	ans, _ := strconv.Atoi(this.tokens[0])

	return ans
}

func evalRPN(tokens []string) int {
	sol := Solution{pos: 0, tokens: tokens}
	return sol.eval()
}
