package main

import (
	"fmt"
	"math"
	"strconv"
)

type Point []int

// FIXME
func (this Point) Distane(other Point) int {
	d := 0
	if this[0] > other[0] {
		d += this[0] - other[0]
	} else {
		d += other[0] - this[0]
	}
	if this[1] > other[1] {
		d += this[1] - other[1]
	} else {
		d += other[1] - this[1]
	}
	return d
}

func (this Point) Id() string {
	return strconv.Itoa(this[0]) + "," + strconv.Itoa(this[1])
}

func filter(points []Point, p Point) []Point {
	remainNuts := []Point{}
	for _, n := range points {
		if n.Id() == p.Id() {
			continue
		}
		remainNuts = append(remainNuts, n)
	}
	return remainNuts
}

type Solution struct {
	height         int
	width          int
	squirrel       Point
	tree           Point
	nuts           []Point
	treeToNuts     map[string]int
	nutsToSquirrel map[string]int
}

func NewSolution(height int, width int, tree []int, squirrel []int, nuts [][]int) *Solution {
	_squirrel := Point(squirrel)
	_tree := Point(tree)
	_nuts := []Point{}
	for _, n := range nuts {
		_nuts = append(_nuts, Point(n))
	}

	treeToNuts := map[string]int{}
	for _, _n := range _nuts {
		treeToNuts[_n.Id()] = _tree.Distane(_n)
	}

	nutsToSquirrel := map[string]int{}
	for _, _n := range _nuts {
		nutsToSquirrel[_n.Id()] = _squirrel.Distane(_n)
	}

	return &Solution{
		height:         height,
		width:          width,
		squirrel:       _squirrel,
		tree:           _tree,
		nuts:           _nuts,
		treeToNuts:     treeToNuts,
		nutsToSquirrel: nutsToSquirrel,
	}
}

func (this *Solution) routeDistance(remainNuts []Point) int {
	if len(remainNuts) == 0 {
		return 0
	}

	min := math.MaxInt64
	for _, nut := range remainNuts {
		// 木からナッツ
		distane := 0
		treeToNut, _ := this.treeToNuts[nut.Id()]
		distane += treeToNut * 2
		md := this.routeDistance(filter(remainNuts, nut))
		if min > md+distane {
			min = md + distane
		}
	}
	return min
}

func (this *Solution) minDistance() int {
	min := math.MaxInt64
	for _, nut := range this.nuts {
		d := 0
		// リスからナッツ
		squirrelToNut, _ := this.nutsToSquirrel[nut.Id()]

		fmt.Printf("リスからナッツ: %d = %s -> %s\n", squirrelToNut, this.squirrel.Id(), nut.Id())

		// ナッツから木
		nutToTree, _ := this.treeToNuts[nut.Id()]

		fmt.Printf("ナッツから木: %d = %s -> %s\n", nutToTree, nut.Id(), this.tree.Id())

		d += squirrelToNut + nutToTree

		fmt.Printf("リスからナッツ、ナッツから木: %d\n", d)

		md := this.routeDistance(filter(this.nuts, nut))

		fmt.Println(md)
		fmt.Println(nut.Id())

		if min > md+d {
			min = md + d
		}
	}

	return min
}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	s := NewSolution(
		height,
		width,
		tree,
		squirrel,
		nuts,
	)
	return s.minDistance()
}
