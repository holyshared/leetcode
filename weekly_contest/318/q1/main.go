package main

/*
非負の整数で構成されるサイズ n の 0 から始まる配列 nums が与えられます。

この配列に n - 1 操作を適用する必要があります。
i 番目の操作 (0 から始まる) では、nums の i 番目の要素に次の操作を適用します。

nums[i] == nums[i + 1] の場合、nums[i] に 2 を掛け、nums[i + 1] を 0 に設定します。
それ以外の場合は、この操作をスキップします。
すべての操作を実行した後、すべての 0 を配列の末尾にシフトします。

たとえば、すべての 0 を最後にシフトした後の配列 [1,0,2,0,0,1] は [1,2,1,0,0,0] です。
結果の配列を返します。

操作は一度にすべてではなく、順番に適用されることに注意してください。

Input: nums = [1,2,2,1,1,0]
Output: [1,4,2,0,0,0]
Explanation: We do the following operations:
- i = 0: nums[0] and nums[1] are not equal, so we skip this operation.
- i = 1: nums[1] and nums[2] are equal, we multiply nums[1] by 2 and change nums[2] to 0. The array becomes [1,4,0,1,1,0].
- i = 2: nums[2] and nums[3] are not equal, so we skip this operation.
- i = 3: nums[3] and nums[4] are equal, we multiply nums[3] by 2 and change nums[4] to 0. The array becomes [1,4,0,2,0,0].
- i = 4: nums[4] and nums[5] are equal, we multiply nums[4] by 2 and change nums[5] to 0. The array becomes [1,4,0,2,0,0].
After that, we shift the 0's to the end, which gives the array [1,4,2,0,0,0].
*/

// 1,2, 2, 1,1,0
// 1 4, 0 , 2,0,0

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			continue
		}
		nums[i] = nums[i] * 2
		nums[i+1] = 0
	}

	i := 0
	ans := make([]int, len(nums))
	for _, num := range nums {
		if num != 0 {
			ans[i] = num
			i++
		}
	}

	return ans
}
