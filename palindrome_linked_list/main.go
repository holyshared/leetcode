package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1a 2a 2b 1b
// slow 2a
// fast 2b
func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode = nil

	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}

	return prev
}

// 1 2 2 1
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	// リストの半分までスキップ(前半)
	// (12345) 54321
	firstHalfEnd := endOfFirstHalf(head)
	// 後半を逆順に並び替える
	// 12345
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart
	result := true

	for result && p2 != nil {
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	// ポインタが変わっているので元に戻してつなぎ直す
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
