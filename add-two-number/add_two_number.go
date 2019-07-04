package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3, present *ListNode
	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2 = l1.Next, l2.Next {
		n := l1.Val + l2.Val + carry
		if n >= 10 {
			carry = 1
			n -= 10
		} else {
			carry = 0
		}
		node := ListNode{n, nil}
		if present == nil {
			l3 = &node
		} else {
			present.Next = &node
		}
		present = &node
	}
	for ; l1 != nil; l1 = l1.Next {
		n := l1.Val + carry
		if n >= 10 {
			carry = 1
			n -= 10
		} else {
			carry = 0
		}
		node := ListNode{n, nil}
		present.Next = &node
		present = &node
	}
	for ; l2 != nil; l2 = l2.Next {
		n := l2.Val + carry
		if n >= 10 {
			carry = 1
			n -= 10
		} else {
			carry = 0
		}
		node := ListNode{n, nil}
		present.Next = &node
		present = &node
	}
	if carry > 0 {
		node := ListNode{carry, nil}
		present.Next = &node
		present = &node
	}
	return l3
}

func main() {
	a := ListNode{9, nil}
	b := ListNode{9, nil}
	c := ListNode{1, nil}
	addTwoNumbers(&a, &c)
}
