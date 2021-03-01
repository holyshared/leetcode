package main

func checkIfExist(arr []int) bool {
	checkedNumbers := map[int]bool{}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			_, ok := checkedNumbers[arr[i]/2]
			if ok {
				return true
			}
		}
		_, ok := checkedNumbers[arr[i]*2]
		if ok {
			return true
		}
		checkedNumbers[arr[i]] = true
	}
	return false
}
