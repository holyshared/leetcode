package main

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
    prevDiff := nums[1] - nums[0]

	// 差がある場合は1、0の場合は1(単一の数値のみ)
	count := 0 
	if prevDiff != 0 {
		count = 2
	} else {
		count = 1
	}

	for i := 2; i < len(nums); i++ {
		diff := nums[i] - nums[i - 1]
		// 前回の差がマイナスかプラスか
		if (diff > 0 && prevDiff <= 0) || (diff < 0 && prevDiff >= 0) {
			count++
			prevDiff = diff
		}
	}
	return count
}
