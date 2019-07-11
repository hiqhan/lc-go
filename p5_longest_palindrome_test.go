package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var s, expect string
	s, expect = "cbbd", "bb"
	if got := LongestPalindrome(s); got != expect {
		t.Errorf("LongestPalindrome(%s) = %s; want %s", s, got, expect)
	}

	s, expect = "abcba", "abcba"
	if got := LongestPalindrome(s); got != expect {
		t.Errorf("LongestPalindrome(%s) = %s; want %s", s, got, expect)
	}

}
