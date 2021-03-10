package main

/*
4 => IV
9 => IX
40 -> XL
90 -> XC
400 -> CD
900 -> CM

I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/

func intToRoman(num int) string {
	remain := num

	roman := ""

	for remain != 0 {
		if remain >= 1000 {
			roman += "M"
			remain = remain - 1000
		} else if remain >= 900 {
			roman += "CM"
			remain = remain - 900
		} else if remain >= 500 {
			roman += "D"
			remain = remain - 500
		} else if remain >= 400 {
			roman += "CD"
			remain = remain - 400
		} else if remain >= 100 {
			roman += "C"
			remain = remain - 100
		} else if remain >= 90 {
			roman += "XC"
			remain = remain - 90
		} else if remain >= 50 {
			roman += "L"
			remain = remain - 50
		} else if remain >= 40 {
			roman += "XL"
			remain = remain - 40
		} else if remain >= 10 {
			roman += "X"
			remain = remain - 10
		} else if remain >= 9 {
			roman += "IX"
			remain = remain - 9
		} else if remain >= 5 {
			roman += "V"
			remain = remain - 5
		} else if remain >= 4 {
			roman += "IV"
			remain = remain - 4
		} else if remain >= 1 {
			roman += "I"
			remain = remain - 1
		}
	}

	return roman
}
