package main

import (
	"math"
)

func _coinChange(coins []int, remain int, memo map[int]int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}
	num, ok := memo[remain-1]
	if ok {
		return num
	}
	min := math.MaxInt64
	for _, c := range coins {
		result := _coinChange(coins, remain-c, memo)
		if result >= 0 && result < min {
			min = 1 + result
		}
	}
	mr := 0
	if min == math.MaxInt64 {
		mr = -1
	} else {
		mr = min
	}
	memo[remain-1] = mr

	return mr
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	memo := map[int]int{}
	_coinChange(coins, amount, memo)

	v, ok := memo[amount-1]
	if ok {
		return v
	} else {
		return -1
	}
}
