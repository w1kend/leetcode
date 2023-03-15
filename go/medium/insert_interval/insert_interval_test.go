package insertinterval

import (
	"leetcode/go/testsupport"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{1, 3},
			},
			want: [][]int{{1, 5}, {6, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{1, 5},
			},
			want: [][]int{{1, 5}, {6, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{1, 6},
			},
			want: [][]int{{1, 7}, {8, 10}, {12, 16}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{17, 20},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}, {17, 20}},
		},
		{
			name: "",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{16, 20},
			},
			want: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 20}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.args.intervals, tt.args.newInterval)
			testsupport.AssertEqual(t, got, tt.want, "unexpected intervals")
		})
	}
}
