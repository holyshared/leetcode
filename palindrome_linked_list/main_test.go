package main

import "testing"

func Test1121(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}

	if isPalindrome(head) != false {
		t.Fatal("oops!!")
	}
}

func Test1234554321(t *testing.T) {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5,
						Next: &ListNode{Val: 5,
							Next: &ListNode{Val: 4,
								Next: &ListNode{Val: 3,
									Next: &ListNode{Val: 2,
										Next: &ListNode{Val: 1}}}}}}}}}}

	if isPalindrome(head) != true {
		t.Fatal("oops!!")
	}
}
