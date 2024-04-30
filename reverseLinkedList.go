package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(list *ListNode) *ListNode {

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

	answer := reverseList(list)
	fmt.Print(answer)
	return
}
