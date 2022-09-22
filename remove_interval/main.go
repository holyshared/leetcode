package main

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	rs, re := toBeRemoved[0], toBeRemoved[1]

	ans := [][]int{}
	for _, interval := range intervals {
		is, ie := interval[0], interval[1]
		if is > re || ie < rs {
			ans = append(ans, []int{is, ie})
		} else {
			if is < rs {
				ans = append(ans, []int{is, rs})
			}
			if ie > re {
				ans = append(ans, []int{re, ie})
			}
		}
	}

	return ans
}
