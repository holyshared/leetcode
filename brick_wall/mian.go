package main

func leastBricks(wall [][]int) int {
	cache := map[int]int{}

	for _, row := range wall {
		sum := 0
		for i := 0; i < len(row)-1; i++ {
			sum += row[i]
			curr, ok := cache[sum]
			if ok {
				cache[sum] = curr + 1
			} else {
				cache[sum] = 1
			}
		}
	}

	res := len(wall)
	keys := []int{}
	for key, _ := range cache {
		keys = append(keys, key)
	}

	for _, key := range keys {
		sum, _ := cache[key]
		if res > len(wall)-sum {
			res = len(wall) - sum
		}
	}
	return res
}
