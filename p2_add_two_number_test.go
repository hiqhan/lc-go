package leetcode

import "testing"

func TestAddTwoNumbers(t *testing.T) {
	a := ListNode{9, nil}
	b := ListNode{9, nil}
	a.Next = &b
	c := ListNode{1, nil}
	expect := string("001")
	if got := addTwoNumbers(&a, &c); got.String() != expect {
		t.Errorf("AddTwoNumbers(%v %v) = %v; want %s", a, c, got, expect)
	}
}
