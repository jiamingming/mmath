package main

import (

	"fmt"
)

func main() {

	n1 := &ListNode{1, nil}
	n2 := &ListNode{4, nil}
	n3 := &ListNode{3, nil}

	n1.Next = n2
	n2.Next = n3

	l1 := n1
	head1 := n1
	for l1 != nil {
		fmt.Print(l1.Val)
		l1 = l1.Next
	}

	m1 := &ListNode{5, nil}
	m2 := &ListNode{6, nil}
	m3 := &ListNode{8, nil}
	m1.Next = m2
	m2.Next = m3
	l2 := m1
	head2 := m1
	for l2 != nil {
		fmt.Print(l2.Val)
		l2 = l2.Next
	}


	res := AddTwoNumbers(head1, head2)
	for res != nil {
		fmt.Print(res.Val)
		res = res.Next
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}


//Input: (1 -> 4 -> 3) + (5 -> 6 -> 8)
//Output: 1-> 2 -> 0 -> 6
//Explanation: 341 + 865 = 1206.

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.Val + l2.Val
	val := sum % 10

	if (l1.Next == nil && l2.Next == nil && sum < 10) {
		return &ListNode{sum, nil}
	}

	if (l1.Next == nil) {
		l1.Next = &ListNode{}
	}

	if (l2.Next == nil) {
		l2.Next = &ListNode{}
	}
	if (sum >= 10) {
		l1.Next.Val += 1
	}
	res := AddTwoNumbers(l1.Next, l2.Next)
	return &ListNode{val, res}
}

