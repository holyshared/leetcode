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

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k := this.storeKey(key)
	items := this.entries[k]
	if items == nil {
		items = []*EntryVal{}
		items = append(items, &EntryVal{Key: key, Val: value})
	} else {
		var entry *EntryVal
		i := 0
		for i = 0; i < len(items); i++ {
			entry = items[i]
			if key != entry.Key {
				continue
			}
		}
		if i > len(items) {
			items = append(items, &EntryVal{Key: key, Val: value})
		} else {
			entry.Val = value
		}
	}
	this.entries[k] = items
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k := this.storeKey(key)
	items := this.entries[k]

	if items == nil {
		return -1
	}

	for _, entry := range items {
		if key != entry.Key {
			continue
		}
		return entry.Val
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k := this.storeKey(key)
	items := this.entries[k]
	if items == nil {
		return
	}

	for i := 0; i < len(items); i++ {
		entry := items[i]
		if key != entry.Key {
			continue
		}
		items = append(items[:i], items[(i+1):]...)
	}
	this.entries[k] = items
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
