package leetcode

import "testing"

func TestIsPalindrome(t *testing.T) {
	var x int
	var expect bool

	x, expect = 121, true
	if got := isPalindrome(x); got != expect {
		t.Errorf("IsPalindrome(%d) = %t; want %t", x, got, expect)
	}
	x, expect = -121, false
	if got := isPalindrome(x); got != expect {
		t.Errorf("IsPalindrome(%d) = %t; want %t", x, got, expect)
	}
	x, expect = 10, false
	if got := isPalindrome(x); got != expect {
		t.Errorf("IsPalindrome(%d) = %t; want %t", x, got, expect)
	}
}
