package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	var a1, a2 []int
	var expect float64
	a1 = []int{10, 20, 30, 40, 50}
	a2 = []int{51, 61, 71, 81, 91}
	expect = 50.5
	if got := FindMedianSortedArrays(a1, a2); got != expect {
		t.Errorf("FindMedianSortedArrays(%v %v) = %f; want %f", a1, a2, got, expect)
	}

	a1 = []int{1, 6, 7, 8}
	a2 = []int{2, 3, 4, 5}
	expect = 4.5
	if got := FindMedianSortedArrays(a1, a2); got != expect {
		t.Errorf("FindMedianSortedArrays(%v %v) = %f; want %f", a1, a2, got, expect)
	}
}
