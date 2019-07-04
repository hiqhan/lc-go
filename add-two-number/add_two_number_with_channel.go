package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func walk(ln *ListNode, ch chan int) {
	var doWalk func(ln *ListNode)
	doWalk = func(ln *ListNode) {
		if ln == nil {
			return
		}
		ch <- ln.Val
		doWalk(ln.Next)
	}
	doWalk(ln)
	close(ch)
}

func Walk(ln *ListNode) string {
	ch := make(chan int)
	go walk(ln, ch)
	var s string
	for {
		i, ok := <-ch
		if ok {
			s += strconv.Itoa(i)
		} else {
			break
		}
	}
	return s
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ch1, ch2 := make(chan int), make(chan int)
	go walk(l1, ch1)
	go walk(l2, ch2)
	var l3, present *ListNode
	carry := 0
	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2
		if !ok1 && !ok2 {
			if carry > 0 {
				node := ListNode{carry, nil}
				present.Next = &node
				present = &node
			}
			break
		}

		if ok1 && ok2 {
			n := n1 + n2 + carry
			if n >= 10 {
				carry = 1
				n = n - 10
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
		} else {
			if ok1 {
				n := n1 + carry
				if n >= 10 {
					carry = 1
					n = n - 10
				} else {
					carry = 0
				}
				node := ListNode{n, nil}
				present.Next = &node
				present = &node
			}
			if ok2 {
				n := n2 + carry
				if n >= 10 {
					carry = 1
					n = n - 10
				} else {
					carry = 0
				}
				node := ListNode{n, nil}
				present.Next = &node
				present = &node
			}
		}
	}
	return l3
}

func main() {
	a := ListNode{9, nil}
	b := ListNode{9, nil}
	c := ListNode{1, nil}
	// b.Next = &c
	a.Next = &b
	fmt.Println(Walk(addTwoNumbers(&a, &c)))
}
