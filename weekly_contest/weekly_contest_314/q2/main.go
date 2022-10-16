package main

/*
サイズ n の整数配列 pref が与えられます。 以下を満たすサイズ n の配列 arr を見つけて返します。

pref[i] = arr[0] ^ arr[1] ^ ... ^ arr[i].
^ はビットごとの xor 演算を表すことに注意してください。

答えが一意であることを証明できます。

- pref[0] = 5.
- pref[1] = 5 ^ 7 = 2.
- pref[2] = 5 ^ 7 ^ 2 = 0.
- pref[3] = 5 ^ 7 ^ 2 ^ 3 = 3.
- pref[4] = 5 ^ 7 ^ 2 ^ 3 ^ 2 = 1.
*/

func findArray(pref []int) []int {
	ans := make([]int, len(pref))

	ans[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		ans[i] = pref[i-1] ^ pref[i]
	}

	return ans
}
