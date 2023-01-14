package main

/*
正の整数 nums の 0 から始まる配列が与えられます。
次の条件を満たすトリプレット (i、j、k) の数を見つけます。

0 <= i < j < k < nums.length
nums[i]、nums[j]、および nums[k] は対ごとに異なります。
つまり、nums[i] != nums[j]、nums[i] != nums[k]、および nums[j] != nums[k] です。
条件を満たすトリプレットの数を返します。
*/

func unequalTriplets(nums []int) int {
	l := len(nums)
	ans := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					ans++
				}
			}
		}
	}

	return ans
}
