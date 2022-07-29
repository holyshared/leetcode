package main

type Vector2D struct {
	vec [][]int
	outer int
	inner int
}

func Constructor(vec [][]int) Vector2D {
	return Vector2D{
		vec: vec,
		outer: 0,
		inner: 0,
	}
}

func (this *Vector2D) Next() int {
	if !this.HasNext() {
		return 0
	}
	curr := this.vec[this.outer][this.inner]
	this.inner++
	return curr
}

func (this *Vector2D) HasNext() bool {
	this.advanceToNext()
	return this.outer < len(this.vec)
}

func (this *Vector2D) advanceToNext() {
	for this.outer < len(this.vec) && this.inner == len(this.vec[this.outer]) {
		this.inner = 0
		this.outer++
	}
}