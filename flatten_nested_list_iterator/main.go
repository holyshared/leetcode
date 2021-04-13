package main

type NestedInteger struct {
	Val   int
	Items []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	if len(this.Items) <= 0 {
		return true
	} else {
		return false
	}
}

func (this NestedInteger) GetInteger() int {
	return this.Val
}
func (n *NestedInteger) SetInteger(value int) {
	n.Val = value
}
func (this *NestedInteger) Add(elem NestedInteger) {
	this.Items = append(this.Items, &elem)
}
func (this NestedInteger) GetList() []*NestedInteger {
	return this.Items
}

func Flatten(items []*NestedInteger, out []int) []int {
	for _, item := range items {
		if item.IsInteger() {
			out = append(out, item.GetInteger())
		} else {
			list := item.GetList()
			out = Flatten(list, out)
		}
	}
	return out
}

type NestedIterator struct {
	items []int
	curr  int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		items: Flatten(nestedList, []int{}),
		curr:  0,
	}
}

func (this *NestedIterator) Next() int {
	val := this.items[this.curr]
	this.curr++
	return val
}

func (this *NestedIterator) HasNext() bool {
	if this.curr > len(this.items)-1 {
		return false
	}
	return true
}
