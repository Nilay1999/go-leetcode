package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicate(head *ListNode) *ListNode {
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

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
		},
	}
	fmt.Print(removeDuplicate(list))
}
