package main

/*
n 個の負でない整数で構成される 0 から始まる配列 nums が与えられます。

1 回の操作で、次のことを行う必要があります。

1 <= i < n および nums[i] > 0 となる整数 i を選択します。
nums[i] を 1 減らします。
nums[i - 1] を 1 増やします。
任意の数の操作を実行した後、nums の最大整数の可能な最小値を返します。


Input: nums = [3,7,1,6]
Output: 5
Explanation:
One set of optimal operations is as follows:
1. Choose i = 1, and nums becomes [4,6,1,6].
2. Choose i = 3, and nums becomes [4,6,2,5].
3. Choose i = 1, and nums becomes [5,5,2,5].
The maximum integer of nums is 5. It can be shown that the maximum number cannot be less than 5.
Therefore, we return 5.

*/

func minimizeArrayValue(nums []int) int {

}
