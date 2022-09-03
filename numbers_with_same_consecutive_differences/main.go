package main

func numsSameConsecDiff(n int, k int) []int {
	queue := []int{}

	for i := 1; i <= 9; i++ {
		queue = append(queue, i)
	}

	for j := 1; j < n; j++ {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			tail := curr % 10
			next := []int{tail + k}

			if k != 0 {
				next = append(next, tail-k)
			}
			for _, v := range next {
				if 0 <= v && v < 10 {
					queue = append(queue, (curr*10)+v)
				}
			}
		}
	}

	return queue
}
