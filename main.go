package main

import (
	"fmt"
	linkedlist "leet/linked_list"
)

func main() {
	l1 := &linkedlist.ListNode{
		Val: 5,
		Next: &linkedlist.ListNode{
			Val: 3,
			Next: &linkedlist.ListNode{
				Val: 0,
			},
		},
	}
	l2 := &linkedlist.ListNode{
		Val: 5,
		Next: &linkedlist.ListNode{
			Val: 1,
		},
	}
	res := linkedlist.AddTwoNumbers(l1, l2)
	fmt.Printf("%#v", res)
}
