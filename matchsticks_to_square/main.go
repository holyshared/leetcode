package main

import "sort"

func makesquare(matchsticks []int) bool {
	sum := 0
	for i := 0; i < len(matchsticks); i++ {
		sum += matchsticks[i]
	}
	if sum%4 != 0 {
		return false
	}
	squareSideLength := sum / 4

	square := make([]int, 4)

	sort.Slice(matchsticks, func (i, j int) bool {
		return matchsticks[i] > matchsticks[j]
	})
	var _makesquare func (curr int) bool
	_makesquare = func (curr int) bool {
		if curr > len(matchsticks)-1 {
			return square[0] == square[1] && square[1] == square[2] && square[2] == square[3]
		}
	
		for i := 0; i < 4; i++ {
			if square[i]+matchsticks[curr] > squareSideLength {
				continue
			}
			square[i] += matchsticks[curr]
	
			if _makesquare(curr+1) {
				return true
			}
			square[i] -= matchsticks[curr]
		}
		return false
	}
	return _makesquare(0)
}
