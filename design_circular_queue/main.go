package main

type MyCircularQueue struct {
	queue []int
	rs    int
	ws    int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k),
		rs:    0,
		ws:    -1,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	nextWs := this.ws + 1
	this.queue[nextWs%len(this.queue)] = value
	this.ws++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.rs++
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.rs%len(this.queue)]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.ws%len(this.queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.ws < this.rs
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.ws-this.rs)+1 == len(this.queue)
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
