package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type List struct {
	head *ListNode
}

func addTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	ls := &ListNode{0, nil}
	temp, carry := ls, 0
	for l1 != nil || l2 != nil {
		a, b, total := 0, 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		total = a + b + carry
		if total > 9 {
			carry = 1
		} else {
			carry = 0
		}
		temp.Next = &ListNode{total % 10, nil}
		temp = temp.Next
	}
	if carry == 1 {
		temp.Next = &ListNode{carry, nil}
	}
	return ls.Next
}

func main() {

	l1 := &ListNode{
		Val:  1,
		Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
	}
	l2 := &ListNode{
		Val:  2,
		Next: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: nil}},
	}
	fmt.Print(addTwoList(l1, l2))
}
