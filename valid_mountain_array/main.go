package main

func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}

	left, right := 0, len(arr)-1

	for left < right-1 {
		if arr[left] < arr[left+1] {
			left++
			continue
		}
		break
	}
	for 0 < right {
		if arr[right] < arr[right-1] {
			right--
			continue
		}
		break
	}

	if right > left || right <= 0 {
		return false
	}
	return true
}
