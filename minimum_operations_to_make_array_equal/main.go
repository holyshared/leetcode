package main

func minOperations(n int) int {
	step := 0
	for n > 0 {
		step += n - 1
		n -= 2
	}
	return step
}
