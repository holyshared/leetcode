package main

import "fmt"


int possibleDups = 0;
int length_ = arr.length - 1; // 7
nums := []int{1,0,2,3,0,4,5,0} 
nums := []int{1,0,2,3,0,4,5,0} // l = 6

// d = 2

// Find the number of zeros to be duplicated
// Stopping when left points beyond the last element in the original array
// which would be part of the modified array
for (int left = 0; left <= length_ - possibleDups; left++) {

		// Count the zeros
		if (arr[left] == 0) {

				// Edge case: This zero can't be duplicated. We have no more space,
				// as left is pointing to the last element which could be included  
				if (left == length_ - possibleDups) {
						// For this zero we just copy it without duplication.
						arr[length_] = 0;
						length_ -= 1;
						break;
				}
				possibleDups++;
		}
}

func duplicateZeros(arr []int) {
	var i int
	clone := make([]int, len(arr))

	for i = 0; i < len(arr); i++ {
		clone[i] = arr[i]
	}

	p := 0

	for i = 0; i < len(arr); i++ {
		if clone[p] == 0 {
			arr[i] = 0
			if i + 1 > len(arr) - 1 {
				return
			}
			arr[i + 1] = 0
			p++
			i++
		} else {
			arr[i] = clone[p]
			p++ //5
		}
	}
}

func main() {
	// 8
	// [1,0,0,2,3,0,0,4]
	nums := []int{1,0,2,3,0,4,5,0} 
	nums := []int{1,0,0,2,3,0,0,4} 
	duplicateZeros(nums);

	fmt.Println(nums)

	nums2 := []int{8,4,5,0,0,0,0,7}
	duplicateZeros(nums2);

	fmt.Println(nums2)
}