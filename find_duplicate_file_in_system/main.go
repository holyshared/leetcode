package main

import "strings"

func pathWithContentFromPathString(path string) [][]string {
	tokens := strings.Split(path, " ")
	directory := tokens[0]
	files := tokens[1:]

	results := make([][]string, len(files))
	for i, file := range files {
		f := strings.Split(file, "(")
		l := len(f[1])
		results[i] = []string{directory + "/" + f[0], f[1][:(l - 1)]}
	}
	return results
}

func findDuplicate(paths []string) [][]string {
	groups := map[string][]string{}

	for _, path := range paths {
		for _, values := range pathWithContentFromPathString(path) {
			path, content := values[0], values[1]
			val, has := groups[content]
			if has {
				groups[content] = append(val, path)
			} else {
				groups[content] = []string{path}
			}
		}
	}

	ans := [][]string{}
	for _, results := range groups {
		if len(results) > 1 {
			ans = append(ans, results)
		}
	}
	return ans
}