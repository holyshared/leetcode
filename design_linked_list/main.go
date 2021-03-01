package main

type Item struct {
	Val  int
	Next *Item
}

type MyLinkedList struct {
	Head *Item
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	curr, ok := this.findAt(index)
	if !ok {
		return -1
	}

	//	if this.Head == nil {
	//	return -1
	//	}

	//	var i int
	//curr := this.Head
	//for i = 0; curr.Next != nil && i < index; i++ {
	//curr = curr.Next
	//	}
	//if curr.Next == nil && i < index {
	//return -1
	//	}

	return curr.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	if this.Head == nil {
		this.Head = &Item{Val: val}
	} else {
		this.Head = &Item{Val: val, Next: this.Head}
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Head == nil {
		this.Head = &Item{Val: val}
		return
	}

	last := this.Head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &Item{Val: val}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		if this.Head == nil {
			return
		}
		this.Head = &Item{Val: val, Next: this.Head}
	} else {
		prev, _ := this.findAt(index - 1)
		if prev == nil {
			return
		}
		next := prev.Next
		prev.Next = &Item{Val: val, Next: next}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 1 2 3 -> 2 3
	if index == 0 {
		if this.Head == nil {
			return
		}
		next, _ := this.findAt(index + 1)
		this.Head = next
		// 1 2 3 -> 1 3
	} else {
		prev, _ := this.findAt(index - 1)
		if prev == nil {
			return
		}
		next, _ := this.findAt(index + 1)
		prev.Next = next
	}
}

func (this *MyLinkedList) findAt(index int) (*Item, bool) {
	if this.Head == nil {
		return nil, false
	}

	var i int
	curr := this.Head
	for i = 0; curr.Next != nil && i < index; i++ {
		curr = curr.Next
	}
	if curr.Next == nil && i < index {
		return nil, false
	}

	return curr, true
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
