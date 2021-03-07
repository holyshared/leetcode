package main

import "testing"

func TestHashMap(t *testing.T) {
	actual := 0

	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	hm.Put(2, 1) // update the existing value

	actual = hm.Get(2) // returns 1
	if actual != 1 {
		t.Fatal("oops!!")
	}

	hm.Remove(2) // remove the mapping for 2

	actual = hm.Get(2)
	if actual != -1 {
		t.Fatal("oops!!")
	}
}

func TestHashMapPut(t *testing.T) {
	actual := 0

	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	hm.Put(3, 3)
	hm.Put(4, 4)
	hm.Put(5, 5)

	puts := []int{1, 2, 3, 4, 5}

	for i, v := range puts {
		actual = hm.Get(i + 1)
		if actual != v {
			t.Fatal("oops!!")
		}
	}
}

func TestHashMapRemove(t *testing.T) {
	actual := 0

	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	hm.Put(3, 3)
	hm.Put(4, 4)
	hm.Put(5, 5)

	hm.Remove(1)
	hm.Remove(2)
	hm.Remove(3)
	hm.Remove(4)
	hm.Remove(5)

	removed := []int{1, 2, 3, 4, 5}

	for _, k := range removed {
		actual = hm.Get(k)
		if actual != -1 {
			t.Fatal("oops!!")
		}
	}
}

func TestHashMapPutNoValue(t *testing.T) {

	hm := Constructor()
	hm.Remove(27)
	hm.Put(65, 65)
	hm.Remove(19)
	hm.Remove(0)
	hm.Get(18)
	hm.Remove(3)
	hm.Put(42, 0)
	hm.Get(19)
	hm.Remove(42)
	hm.Put(17, 90)
	hm.Put(31, 76)
	hm.Put(48, 71)
	hm.Put(5, 50)
	hm.Put(7, 68)
	hm.Put(73, 74)
	hm.Put(85, 18)
	hm.Put(74, 95)
	hm.Put(84, 82)
	hm.Put(59, 29)
	hm.Put(71, 71)
	hm.Remove(42)
	hm.Put(51, 40)
	hm.Put(33, 76)
	hm.Get(17)
	hm.Put(89, 95)
	hm.Get(95)
	hm.Put(30, 31)
	hm.Put(37, 99)
	hm.Get(51)
	hm.Put(95, 35)
	hm.Remove(65)
	hm.Remove(81)
	hm.Put(61, 46)
	hm.Put(50, 33)
	hm.Get(59)
	hm.Remove(5)
	hm.Put(75, 89)
	hm.Put(80, 17)
	hm.Put(35, 94)
	hm.Get(80)
	hm.Put(19, 68)
	hm.Put(13, 17)
	hm.Remove(70)
	hm.Put(28, 35)
	hm.Remove(99)
	hm.Remove(37)
	hm.Remove(13)
	hm.Put(90, 83)
	hm.Remove(41)
	hm.Get(50)
	hm.Put(29, 98)
	hm.Put(54, 72)
	hm.Put(6, 8)
	hm.Put(51, 88)
	hm.Remove(13)
	hm.Put(8, 22)
	hm.Get(85)
	hm.Put(31, 22)
	hm.Put(60, 9)
	hm.Get(96)
	hm.Put(6, 35)
	hm.Remove(54)
	hm.Get(15)
	hm.Get(28)
	hm.Remove(51)
	hm.Put(80, 69)
	hm.Put(58, 92)
	hm.Put(13, 12)
	hm.Put(91, 56)
	hm.Put(83, 52)
	hm.Put(8, 48)
	hm.Get(62)
	hm.Get(54)
	hm.Remove(25)
	hm.Put(36, 4)
	hm.Put(67, 68)
	hm.Put(83, 36)
	hm.Put(47, 58)
	hm.Get(82)
	hm.Remove(36)
	hm.Put(30, 85)
	hm.Put(33, 87)
	hm.Put(42, 18)
	hm.Put(68, 83)
	hm.Put(50, 53)
	hm.Put(32, 78)
	hm.Put(48, 90)
	hm.Put(97, 95)
	hm.Put(13, 8)
	hm.Put(15, 7)
	hm.Remove(5)
	hm.Remove(42)
	hm.Get(20)
	hm.Remove(65)
	hm.Put(57, 9)
	hm.Put(2, 41)
	hm.Remove(6)
	hm.Get(33)
	hm.Put(16, 44)
	hm.Put(95, 30)

}
