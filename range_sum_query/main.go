package main

type NumArray struct {
	nums []int
	tree []int
}

func buildTree(nums []int) []int {
	length := len(nums)
	tree := make([]int, length*2)
	// [0, 0, 0, 1, 3, 5]
	for i, j := length, 0; i < 2*length; i++ {
		tree[i] = nums[j]
		j++
	}
	// [9, 9, 8, 1, 3, 5]
	for i := length - 1; i > 0; i-- {
		tree[i] = tree[i*2] + tree[i*2+1]
	}
	return tree
}

func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
		tree: buildTree(nums),
	}
}

func (this *NumArray) Update(index int, val int) {
	// 指定された場所の値を置き換える
	index += len(this.nums)
	this.tree[index] = val

	// 合計値を再計算し直す
	// [9, 9, 8, 1, 3, 5]
	for index > 0 {
		left := index
		right := index
		if index%2 == 0 {
			right = index + 1
		} else {
			left = index - 1
		}
		this.tree[index/2] = this.tree[left] + this.tree[right]
		index /= 2
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	left += len(this.nums)
	right += len(this.nums)
	sum := 0
	for left <= right {
		if (left % 2) == 1 {
			sum += this.tree[left]
			left++
		}
		if (right % 2) == 0 {
			sum += this.tree[right]
			right--
		}
		left /= 2
		right /= 2
	}
	return sum
}
