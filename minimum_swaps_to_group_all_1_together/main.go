package main

func minSwaps(data []int) int {
	ones := 0
	for i := 0; i < len(data); i++ {
		ones += data[i]
	}

	oneCount, maxOne := 0, 0
	left, right := 0, 0

	for right < len(data) {
		oneCount += data[right]
		right++
		if right-left > ones {
			oneCount -= data[left]
			left++
		}
		if maxOne < oneCount {
			maxOne = oneCount
		}
	}
	return ones - maxOne
}
