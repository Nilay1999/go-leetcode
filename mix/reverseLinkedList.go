package mix

func ReverseList(list *ListNode) *ListNode {

	ptr := list
	prev := &ListNode{Val: 0, Next: nil}

	for ptr != nil {
		temp := ptr.Next
		ptr.Next = prev
		prev = ptr
		ptr = temp
	}

	return prev
}
