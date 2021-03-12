package main

var chars = []string{"0", "1"}

func pattern(chars []string, k int, paths []string, results [][]string) [][]string {
	if k == 0 {
		fixed := make([]string, len(paths))
		copy(fixed, paths)
		results = append(results, fixed)
		return results
	}
	for _, v := range chars {
		paths = append(paths, v)
		results = pattern(chars, k-1, paths, results)
		paths = paths[:(len(paths) - 1)]
	}
	return results
}

func hasAllCodes(s string, k int) bool {
	results := [][]string{}
	results = pattern(chars, k, []string{}, results)
	return true
}
