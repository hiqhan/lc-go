package leetcode

import "testing"

func TestReverse(t *testing.T) {
	var i, expect int
	i, expect = 123, 321
	if got := reverse(i); got != expect {
		t.Errorf("Reverse(%d) = %d; want %d", i, got, expect)
	}

	i, expect = -123, -321
	if got := reverse(i); got != expect {
		t.Errorf("Reverse(%d) = %d; want %d", i, got, expect)
	}

	i, expect = 1200, 21
	if got := reverse(i); got != expect {
		t.Errorf("Reverse(%d) = %d; want %d", i, got, expect)
	}

	i, expect = 1534236469, 0
	if got := reverse(i); got != expect {
		t.Errorf("Reverse(%d) = %d; want %d", i, got, expect)
	}

	i, expect = 0, 0
	if got := reverse(i); got != expect {
		t.Errorf("Reverse(%d) = %d; want %d", i, got, expect)
	}
}
