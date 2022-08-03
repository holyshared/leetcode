package main

import (
	"github.com/emirpasic/gods/maps/treemap"
)

type MyCalendar struct {
	calendar *treemap.Map
}

func Constructor() MyCalendar {
	return MyCalendar{
		calendar: treemap.NewWithIntComparator(),
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	prev, _ := this.calendar.Floor(start)
	next, _ := this.calendar.Ceiling(start)

	pass := false
	if prev != nil {
		value, _ := this.calendar.Get(prev)
		pass = prev == nil || value.(int) <= start
	} else {
		pass = true
	}

	if pass &&
		(next == nil || end <= next.(int)) {
		this.calendar.Put(start, end)
		return true
	}
	return false
}
