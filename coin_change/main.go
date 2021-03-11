package main

import (
	"fmt"
	"math"
)

// memo[x]は残額をキーとしたコインの最小枚数
// キーが残額で、値がコインの枚数
func _coinChange(coins []int, remain int, memo []int) int {
	if remain < 0 {
		return -1
	}
	if remain == 0 {
		return 0
	}
	// すでに残額からの最小枚数がある場合それを返す
	// 最小枚数かどうかは補償されていない
	// 他の消費パターンで更新される可能性があるため
	if memo[remain-1] != 0 {
		return memo[remain-1]
	}
	min := math.MaxInt64
	for _, c := range coins {
		result := _coinChange(coins, remain-c, memo)
		// コインを消費できた && 現在の最小値より少ない枚数?
		if result >= 0 && result < min {
			min = 1 + result // コインの枚数を増やす
		}
	}
	mr := 0
	// 手持ちのコインだと0にできない
	if min == math.MaxInt64 {
		mr = -1
	} else {
		mr = min
	}
	// 今の残額での最小枚数
	memo[remain-1] = mr

	return mr
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	bucket := make([]int, amount)
	result := _coinChange(coins, amount, bucket)

	fmt.Println(bucket)

	return result
}
