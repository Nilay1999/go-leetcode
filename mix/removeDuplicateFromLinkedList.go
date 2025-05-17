package mix

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveDuplicate(head *ListNode) *ListNode {
	temp := &ListNode{Next: head}
	prev := &ListNode{0, nil}
	prev = temp
	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			prev.Next = head.Next
		} else {
			prev = prev.Next
		}
		head = head.Next
	}
	return temp.Next
}
