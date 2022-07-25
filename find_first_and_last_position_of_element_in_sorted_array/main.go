package main

func findBound(nums []int, target int, isFirst bool) int {
	begin := 0
  end := len(nums) - 1

	for begin <= end {
		mid := (begin + end) / 2
    if nums[mid] == target {
			// 開始を検索する場合
			if isFirst {
				if mid == begin || nums[mid - 1] < target {
        	return mid
        }
        end = mid - 1
			// 終了を検索する場合
      } else {
				if mid == end || nums[mid + 1] > target {
        	return mid
        }
				begin = mid + 1
      }
    } else if nums[mid] > target {
			end = mid - 1
    } else {
			begin = mid + 1
    }
  }
  return -1
}

func searchRange(nums []int, target int) []int {
  lowerBound := findBound(nums, target, true)
  if lowerBound == -1 {
    return []int{-1, -1}
  }
  upperBound := findBound(nums, target, false)

  return []int{lowerBound, upperBound}
}
