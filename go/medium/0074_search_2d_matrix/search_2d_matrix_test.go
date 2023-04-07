package search2dmatrix

import (
	"leetcode/go/test"
	"testing"
)

func Test_searchMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
		{
			name: "e2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "e3",
			args: args{
				matrix: [][]int{{1}},
				target: 0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix(tt.args.matrix, tt.args.target)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
