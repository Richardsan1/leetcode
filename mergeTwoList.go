package main

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{0, nil}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return dummy.Next
}

// func main() {
// 	// list1 = [1,2,4], list2 = [1,3,4]
// 	n3 := &ListNode{4, nil}
// 	n2 := &ListNode{2, n3}
// 	list1 := &ListNode{1, n2}

// 	n5 := &ListNode{4, nil}
// 	n4 := &ListNode{3, n5}
// 	list2 := &ListNode{1, n4}

// 	mergeTwoLists(list1, list2)
// }
