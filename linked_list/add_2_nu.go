package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func run() {
	l1 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 0,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 1,
		},
	}
	res := AddTwoNumbers(l1, l2)
	fmt.Printf("%#v", res)
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tail *ListNode
	carry := 0
	for {
		sum := 0 + carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l1.Val
			l2 = l2.Next
		}
		if sum > 9 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		if result == nil {
			result = new(ListNode)
			tail = result
		} else {
			tail.Next = new(ListNode)
			tail = tail.Next
		}
		tail.Val = sum
		if l1 == nil && l2 == nil {
			if carry != 0 {
				tail.Next = new(ListNode)
				tail.Next.Val = carry
			}
			return result
		}
	}
}

func rec(l1 *ListNode, l2 *ListNode) (*ListNode, int) {
	if l1 == nil && l2 == nil {
		return nil, 0
	}
	if l1 == nil {
		return rec(l1, l2.Next)
	} else if l2 == nil {
		return rec(l1.Next, l2)
	}

	result, carry := rec(l1.Next, l2.Next)
	val := l1.Val + l2.Val + carry
	carry = 0
	if val > 9 {
		carry = 1
		val = val - 10
	}
	return &ListNode{
		Val:  val,
		Next: result,
	}, carry
}
