package main

type EntryVal struct {
	Key int
	Val int
}

type MyHashMap struct {
	entries [][]*EntryVal
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{entries: make([][]*EntryVal, 1000000)}
}

func (this *MyHashMap) storeKey(key int) int {
	return key % 1000000
}

func (this *MyHashMap) getEntries(key int) []*EntryVal {
	k := this.storeKey(key)
	items := this.entries[k]
	if items == nil {
		return []*EntryVal{}
	}
	return items
}

func (this *MyHashMap) findEntry(key int) (*EntryVal, bool) {
	items := this.getEntries(key)
	if len(items) <= 0 {
		return nil, false
	}

	var entry *EntryVal
	for _, entry = range items {
		if key == entry.Key {
			return entry, true
		}
	}

	return nil, false
}

func (this *MyHashMap) findEntryIndex(key int) (int, bool) {
	items := this.getEntries(key)
	if len(items) <= 0 {
		return -1, false
	}

	for i, entry := range items {
		if key == entry.Key {
			return i, true
		}
	}

	return -1, false
}

func (this *MyHashMap) addEntry(key int, value int) {
	items := this.getEntries(key)
	this.entries[this.storeKey(key)] = append(items, &EntryVal{Key: key, Val: value})
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	entry, ok := this.findEntry(key)
	if ok {
		entry.Val = value
		return
	}
	this.addEntry(key, value)
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	entry, ok := this.findEntry(key)
	if !ok {
		return -1
	}
	return entry.Val
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	i, ok := this.findEntryIndex(key)
	if !ok {
		return
	}
	items := this.getEntries(key)
	this.entries[this.storeKey(key)] = append(items[:i], items[(i+1):]...)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
