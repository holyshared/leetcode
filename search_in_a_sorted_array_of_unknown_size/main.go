package main

type ArrayReader struct {
	Values []int
}

func (this *ArrayReader) get(index int) int {
	if len(this.Values)-1 >= index {
		return this.Values[index]
	}
	return 2147483647
}

func search(reader ArrayReader, target int) int {
	if reader.get(0) == target {
		return 0
	}

	left := 0
	right := 1

	for reader.get(right) < target {
		left = right
		right <<= 1
	}

	var pivot, num int
	for left <= right {
		pivot = left + ((right - left) >> 1)
		num = reader.get(pivot)

		if num == target {
			return pivot
		}
		if num > target {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return -1
}
