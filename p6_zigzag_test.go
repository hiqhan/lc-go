package leetcode

import "testing"

func TestConvert(t *testing.T) {
	var s, cs string
	var n int
	s, n, cs = string("PAYPALISHIRING"), 3, string("PAHNAPLSIIGYIR")
	if got := convert(s, n); got != cs {
		t.Errorf("Convert(%s, %d) = %s; want %s", s, n, got, cs)
	}

	s, n, cs = string("PAYPALISHIRING"), 4, string("PINALSIGYAHRPI")
	if got := convert(s, n); got != cs {
		t.Errorf("Convert(%s, %d) = %s; want %s", s, n, got, cs)
	}

	s, n, cs = string("PAYPALISHIRING"), 1, string("PAYPALISHIRING")
	if got := convert(s, n); got != cs {
		t.Errorf("Convert(%s, %d) = %s; want %s", s, n, got, cs)
	}
}

// func BenchmarkConvert(b *testing.B) {
// 	s, n := string("PAYPALISHIRING"), 3
// 	for i := 0; i < b.N; i++ {
// 		Convert(s, n)
// 	}
// }
