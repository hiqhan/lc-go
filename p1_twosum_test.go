package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums, target := []int{1, 3, 3}, 6
	expect := []int{1, 2}
	if got := TwoSum(nums, target); !reflect.DeepEqual(got, expect) {
		t.Errorf("TwoSum(%v, %d) = %v; want %v", nums, target, got, expect)
	}
}
