package main

func garbageCollection(garbage []string, travel []int) int {
	sums := make([]int, 3)

	travelMinutes := make([]int, len(travel)+1)
	for i := 0; i < len(travel); i++ {
		travelMinutes[i+1] = travelMinutes[i] + travel[i]
	}

	collectMinutes := 0

	for i := 0; i < len(garbage); i++ {
		for _, v := range garbage[i] {
			ty := 0
			if v == 'G' {
				ty = 0
			} else if v == 'P' {
				ty = 1
			} else {
				ty = 2
			}
			if i > 0 {
				sums[ty] = travelMinutes[i]
			}
			collectMinutes++
		}
	}

	for _, v := range sums {
		collectMinutes += v
	}

	return collectMinutes
}
