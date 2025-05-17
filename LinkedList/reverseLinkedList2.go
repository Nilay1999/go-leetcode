package main

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	temp := &ListNode{Next: head}
	prev := temp
	ptr := &ListNode{}
	ptr = head
	l := 1

	for l != left {
		ptr = ptr.Next
		prev = prev.Next
		l += 1
	}

	return head
}
