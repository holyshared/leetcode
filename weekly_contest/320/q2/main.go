package main

/*
二分探索木のルートと、正の整数で構成されるサイズ n の配列クエリが与えられます。

answer[i] = [mini, maxi] であるサイズ n の 2D 配列の答えを見つけます。

mini は、queries[i] 以下のツリー内の最大値です。 そのような値が存在しない場合は、代わりに -1 を追加します。
maxi は、queries[i] 以上のツリー内の最小値です。 そのような値が存在しない場合は、代わりに -1 を追加します。
配列の答えを返します。

入力: ルート = [6,2,13,1,4,9,15,null,null,null,null,null,null,14]、クエリ = [2,5,16]
出力: [[2,2],[4,6],[15,-1]]
説明: 次の方法でクエリに回答します。
- ツリー内の 2 以下の最大数は 2 であり、2 以上の最小数は依然として 2 です。したがって、最初のクエリの答えは [2,2] です。
- ツリー内の 5 以下の最大数は 4 で、5 以上の最小数は 6 です。したがって、2 番目のクエリの答えは [4,6] です。
- ツリー内の 16 以下の最大数は 15 であり、16 以上の最小数は存在しません。 したがって、3 番目のクエリの答えは [15,-1] です。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solution struct {
	values []int
}

func (this *Solution) nodeToSlice(node *TreeNode) {
	if node == nil {
		return
	}
	this.nodeToSlice(node.Left)
	this.values = append(this.values, node.Val)
	this.nodeToSlice(node.Right)
}

func (this *Solution) finMin(q int) int {
	left, right, ans := 0, len(this.values)-1, -1

	for left <= right {
		mid := left + (right-left)/2
		if this.values[mid] <= q {
			if this.values[mid] == q {
				ans = q
				break
			}
			ans = this.values[mid]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func (this *Solution) finMax(q int) int {
	left, right, ans := 0, len(this.values)-1, -1

	for left <= right {
		mid := left + (right-left)/2
		if this.values[mid] >= q {
			if this.values[mid] == q {
				ans = q
				break
			}
			ans = this.values[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func closestNodes(root *TreeNode, queries []int) [][]int {
	ans := [][]int{}

	sol := &Solution{values: []int{}}
	sol.nodeToSlice(root)

	for _, q := range queries {
		minv := sol.finMin(q)
		maxv := sol.finMax(q)
		ans = append(ans, []int{minv, maxv})
	}
	return ans
}
