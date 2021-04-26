package main

func _furthestBuilding(heights []int, curr int, remainBricks int, remainLadders int) int {
	if curr >= len(heights)-1 {
		return curr
	}

	if heights[curr] >= heights[curr+1] {
		return _furthestBuilding(heights, curr+1, remainBricks, remainLadders)
	}

	need := heights[curr+1] - heights[curr]
	// レンガは使える
	canUseBricks := need < remainBricks
	b := curr
	if canUseBricks {
		b = _furthestBuilding(heights, curr+1, remainBricks-need, remainLadders)
	}
	canUseRemainLadders := remainLadders > 0
	l := curr
	if canUseRemainLadders {
		l = _furthestBuilding(heights, curr+1, remainBricks, remainLadders-1)
	}

	if l < b {
		return b
	} else {
		return l
	}
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	return _furthestBuilding(heights, 0, bricks, ladders)
}
