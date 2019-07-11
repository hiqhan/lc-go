package leetcode

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	var s string
	var expect int

	s, expect = "museuwzbczdejna", 8
	if got := LengthOfLongestSubstring(s); got != expect {
		t.Errorf("LengthOfLongestSubstring(%s) = %d; want %d", s, got, expect)
	}

	s, expect = "bbbb", 1
	if got := LengthOfLongestSubstring(s); got != expect {
		t.Errorf("LengthOfLongestSubstring(%s) = %d; want %d", s, got, expect)
	}
}
