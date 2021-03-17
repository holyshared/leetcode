package main

import (
	"math"
	"math/rand"
)

type Solution struct {
	radius float64
	xc     float64
	yc     float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		radius: radius,
		xc:     x_center,
		yc:     y_center,
	}
}

func (this *Solution) RandPoint() []float64 {
	x0, y0 := this.xc-this.radius, this.yc-this.radius

	for {
		xg := x0 + rand.Float64()*this.radius*2
		yg := y0 + rand.Float64()*this.radius*2

		// 2点間の距離を求める
		// 円の中心から、ランダムで生成した点の距離
		a := math.Pow((xg - this.xc), 2) // p1
		b := math.Pow((yg - this.yc), 2) // p2
		d := math.Sqrt(a + b)

		// 2点間の距離が半径以下の場合は円内に収まっている
		if d <= this.radius {
			return []float64{xg, yg}
		}
	}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
