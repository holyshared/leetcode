package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(src int, letters [][]rune, visited []int, component []int, minRune int) ([]int, int) {
	visited[src] = 1
	component = append(component, src)

	minRune = min(minRune, src)

	for i := 0; i < 26; i++ {
		if letters[src][i] == 1 && visited[i] == 0 {
			component, minRune = dfs(i, letters, visited, component, minRune)
		}
	}
	return component, minRune
}

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	letters := make([][]rune, 26)
	for i, _ := range letters {
		letters[i] = make([]rune, 26)
	}

	s1r, s2r := []rune(s1), []rune(s2)
	for i := 0; i < len(s1r); i++ {
		letters[s1r[i]-'a'][s2r[i]-'a'] = 1
		letters[s2r[i]-'a'][s1r[i]-'a'] = 1
	}

	mapping := make([]rune, 26)
	for i := 0; i < 26; i++ {
		mapping[i] = rune(i) + 'a'
	}

	visited := make([]int, 26)
	for c := 0; c < 26; c++ {
		if visited[c] == 0 {
			component, minRune := dfs(c, letters, visited, []int{}, 27)
			for _, vertex := range component {
				mapping[vertex] = rune(minRune)
			}
		}
	}

	ans := []rune{}

	for _, c := range []rune(baseStr) {
		ans = append(ans, mapping[c-'a']+'a')
	}

	return string(ans)
}
