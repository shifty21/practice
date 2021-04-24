package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := ListNode{}
	dummy.Next = nil
	tail := &dummy

	merge(tail, l1, l2)
	return tail.Next
}

func merge(result *ListNode, l1, l2 *ListNode) {
	if l1 == nil {
		(*result).Next = l2
		return
	}
	if l2 == nil {
		(*result).Next = l1
		return
	}
	if l1.Val >= l2.Val {
		(*result).Next = l2
		l2 = (*l2).Next
		result = (*result).Next
		merge(result, l1, l2)
	} else {
		(*result).Next = l1
		l1 = (*l1).Next
		result = (*result).Next
		merge(result, l1, l2)
	}
}
