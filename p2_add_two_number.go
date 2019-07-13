package leetcode

import "strconv"

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

func str(ln *ListNode) string {
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
	dummyHead := new(ListNode)
	p, q, curr := l1, l2, dummyHead
	carry := 0
	for p != nil || q != nil {
		var x, y int
		if x = 0; p != nil {
			x = p.Val
		}
		if y = 0; q != nil {
			y = q.Val
		}
		sum := x + y + carry
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{1, nil}
	}
	return dummyHead.Next
}

func (l *ListNode) String() string {
	return str(l)
}
