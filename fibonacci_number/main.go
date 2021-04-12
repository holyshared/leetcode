package main

func _fib(n int, cache map[int]int) (int, map[int]int) {
	cacheResult, ok := cache[n]
	if ok {
		return cacheResult, cache
	}

	result := 0
	if n < 2 {
		result = n
	} else {
		n1, cache := _fib(n-1, cache)
		n2, cache := _fib(n-2, cache)
		result = n1 + n2
	}
	cache[n] = result
	return result, cache
}

func fib(n int) int {
	result, _ := _fib(n, map[int]int{})
	return result
}
