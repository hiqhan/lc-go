package leetcode

import "testing"

func Test_solveSudoku(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSudoku(tt.args.board); got != tt.want {
				t.Errorf("solveSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
