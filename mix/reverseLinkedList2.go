package mix

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	temp := &ListNode{Next: head}
	prev := &ListNode{0, nil}
	prev = temp
	ptr := &ListNode{0, nil}
	ptr = head
	l := 1

	for l != left {
		ptr = ptr.Next
		prev = prev.Next
		l += 1
	}

	return head
}
