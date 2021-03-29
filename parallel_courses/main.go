package main

func _tryCompelete(dep map[int][]int, completed map[int]bool) (map[int][]int, map[int]bool) {
	deleted := []int{}

	for dk, deps := range dep {
		compok := true
		for _, depc := range deps {
			_, cmp := completed[depc]
			if !cmp {
				compok = false
				break
			}
		}
		if compok {
			deleted = append(deleted, dk)
		}
	}

	for _, v := range deleted {
		completed[v] = true
		delete(dep, v)
	}

	return dep, completed
}

func minimumSemesters(n int, relations [][]int) int {
	depIndex := map[int][]int{}

	for i := 0; i < len(relations); i++ {
		rel := relations[i]
		dep, ok := depIndex[rel[1]]
		if ok {
			depIndex[rel[1]] = append(dep, rel[0])
		} else {
			depIndex[rel[1]] = []int{rel[0]}
		}
	}

	completed := map[int]bool{}
	for i := 1; i < n; i++ {
		_, ok := depIndex[i]
		if ok {
			continue
		}
		completed[i] = true
	}

	level := 1
	for len(depIndex) > 0 {
		depC := len(depIndex)
		depIndex, completed = _tryCompelete(depIndex, completed)
		if depC > len(depIndex) {
			level++
		} else {
			return -1
		}
	}
	return level
}
