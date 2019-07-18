package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {
	var s string
	var expect int

	s, expect = "42", 42
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "   -42", -42
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "   -4   2", -4
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "4193 with words", 4193
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "words and 987", 0
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "-91283472332", -2147483648
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "91283472332", 2147483647
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
	s, expect = "0-1", 0
	if got := myAtoi(s); got != expect {
		t.Errorf("myAtoi(%s) = %d; want %d", s, got, expect)
	}
}
