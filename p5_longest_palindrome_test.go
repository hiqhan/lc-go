package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var s, expect string
	s, expect = "cbbd", "bb"
	if got := longestPalindrome(s); got != expect {
		t.Errorf("LongestPalindrome(%s) = %s; want %s", s, got, expect)
	}

	s, expect = "abcba", "abcba"
	if got := longestPalindrome(s); got != expect {
		t.Errorf("LongestPalindrome(%s) = %s; want %s", s, got, expect)
	}

}
