package main

// A = 1, 2, 0
/*
l:
1 2
2 0 *

g:
1 2
1 0 *
2 0 *
*/

func isIdealPermutation(A []int) bool {
	N := len(A)
	floor := N

	for i := N - 1; i >= 2; i-- {
		if floor > A[i] {
			floor = A[i]
		}
		if A[i-2] > floor {
			return false
		}
	}
	return true
}
