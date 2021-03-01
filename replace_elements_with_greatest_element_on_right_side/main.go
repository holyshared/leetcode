package main

func findMax(arr []int) int {
	maxi := 0
	for i := 0; i < len(arr); i++ {
		if arr[maxi] >= arr[i] {
			continue
		}
		maxi = i
	}
	return maxi
}

func replaceElements(arr []int) []int {
	max := -1

	for i := 0; i < len(arr)-1; i++ {
		if max < 0 {
			max = findMax(arr[(i + 1):])
			arr[i] = arr[max+i+1]
			max = max + i + 1
		} else if i+1 <= max {
			arr[i] = arr[max]
		} else {
			max = findMax(arr[(i + 1):])
			arr[i] = arr[max+i+1]
			max = max + i + 1
		}
	}
	arr[len(arr)-1] = -1

	return arr
}
