package main

import "sort"

func scheduleCourse(courses [][]int) int {
	sort.SliceStable(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	time, count := 0, 0

	for i := 0; i < len(courses); i++ {
		if time+courses[i][0] <= courses[i][1] {
			time += courses[i][0]
			count++
		} else {
			max_i := i
			for j := 0; j < i; j++ {
				if courses[j][0] > courses[max_i][0] {
					max_i = j
				}
			}
			if courses[max_i][0] > courses[i][0] {
				time += courses[i][0] - courses[max_i][0]
			}
			courses[max_i][0] = -1
		}
	}
	return count
}
