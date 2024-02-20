package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum int
	var carry int

	var firstNode *ListNode = &ListNode{0, nil}
	var dummy *ListNode = firstNode
	for l1 != nil || l2 != nil || carry != 0 {
		var digit1, digit2 int
		if l1 != nil {
			digit1 = l1.Val
			l1 = l1.Next
		} else {
			digit1 = 0
		}
		if l2 != nil {
			digit2 = l2.Val
			l2 = l2.Next
		} else {
			digit2 = 0
		}

		sum = digit1 + digit2 + carry
		digit := sum % 10
		carry = sum / 10

		newNode := &ListNode{digit, nil}
		dummy.Next = newNode
		dummy = dummy.Next
	}

	return firstNode.Next
}
