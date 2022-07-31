package main

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

func (this *NumArray) Update(index int, val int)  {
	this.nums[index] = val
}

func (this *NumArray) SumRange(left int, right int) int {
	if left == right {
		return this.nums[left]
	}

	sum := 0

	for left < right {
		sum += this.nums[left]
		sum += this.nums[right]
		left++
		right--
	}

	if left == right {
		sum += this.nums[left]
	}

	return sum
}
