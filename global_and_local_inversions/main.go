package main

func isIdealPermutation(A []int) bool {
	N := len(A)

	l := 0
	for i := 0; i < N-1; i++ {
		if A[i] > A[i+1] {
			l++
		}
	}

	g := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if A[i] > A[j] {
				g++
			}
		}
	}

	return l == g
}
