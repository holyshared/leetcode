package main

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	rs, re := toBeRemoved[0], toBeRemoved[1]

	ans := [][]int{}
	for _, interval := range intervals {
		is, ie := interval[0], interval[1]
		if is >= rs && ie <= re {
			continue
		} else if ie > rs && ie < re || is < rs && ie == re {
			ans = append(ans, []int{is, rs})
		} else if rs < is && re > ie || is < re && re < ie && rs < is || rs == is && re < ie {
			ans = append(ans, []int{re, ie})
		} else if rs > is && re < ie {
			ans = append(ans, []int{is, rs})
			ans = append(ans, []int{re, ie})
		} else {
			ans = append(ans, []int{is, ie})
		}
	}

	return ans
}
