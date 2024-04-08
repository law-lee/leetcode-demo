package main

type ListNode struct {
	val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	pre := &ListNode{}

	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}

func help(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return help(head, next)
}
