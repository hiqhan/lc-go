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
		{
			name: "frist test",
			args: args{
				[][]byte{
					// {5, 3, e, e, 7, e, e, e, e},
					// {6, e, e, 1, 9, 5, e, e, e},
					// {e, 9, 8, e, e, e, e, 6, e},
					// {8, e, e, e, 6, e, e, e, 3},
					// {4, e, e, 8, e, 3, e, e, 1},
					// {7, e, e, e, 2, e, e, e, 6},
					// {e, 6, e, e, e, e, 2, 8, e},
					// {e, e, e, 4, 1, 9, e, e, 5},
					// {e, e, e, e, 8, e, e, 7, 9},
					[]byte(string("53..7....")),
					[]byte(string("6..195...")),
					[]byte(string(".98....6.")),
					[]byte(string("8...6...3")),
					[]byte(string("4..8.3..1")),
					[]byte(string("7...2...6")),
					[]byte(string(".6....28.")),
					[]byte(string("...419..5")),
					[]byte(string("....8..79")),
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveSudoku(tt.args.board); got != tt.want {
				t.Errorf("solveSudoku() = %v, want %v", got, tt.want)
			}
		})
	}
}
