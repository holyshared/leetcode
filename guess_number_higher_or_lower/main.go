package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guess(num int) int {
	if num > 6 {
		return -1
	} else if num < 6 {
		return 1
	} else {
		return 0
	}
}

func _guessNumber(left int, right int) int {
	mid := (left + right) / 2
	if guess(mid) == 1 {
		return _guessNumber(mid+1, right)
	} else if guess(mid) == -1 {
		return _guessNumber(left, mid-1)
	} else {
		return mid
	}
}

func guessNumber(n int) int {
	return _guessNumber(1, n)
}
