package main

func distributeCandies(candyType []int) int {
	types := map[int]bool{}

	for i := 0; i < len(candyType); i++ {
		types[candyType[i]] = true
	}
	eatCandyCount := len(candyType) / 2

	if eatCandyCount < len(types) {
		return eatCandyCount
	} else {
		return len(types)
	}
}
