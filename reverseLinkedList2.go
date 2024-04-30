package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
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
	for ptr != nil {

	}

	return head
}

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val:  6,
							Next: nil,
						},
					},
				},
			},
		},
	}
	left, right := 2, 4
	fmt.Print(reverseBetween(list, left, right))
}
