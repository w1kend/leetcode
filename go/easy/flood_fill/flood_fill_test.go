package flood_fill

import (
	"leetcode/go/test"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		image [][]int
		sr    int
		sc    int
		color int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:    1,
				sc:    1,
				color: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			name: "",
			args: args{
				image: [][]int{
					{1, 0, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:    1,
				sc:    1,
				color: 3,
			},
			want: [][]int{
				{3, 0, 1},
				{3, 3, 0},
				{3, 0, 1},
			},
		},
		{
			name: "",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
				sr:    1,
				sc:    0,
				color: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
		},
		{
			name: "",
			args: args{
				image: [][]int{
					{0, 0, 0},
					{1, 0, 0},
				},
				sr:    1,
				sc:    0,
				color: 2,
			},
			want: [][]int{
				{0, 0, 0},
				{2, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := floodFill(tt.args.image, tt.args.sr, tt.args.sc, tt.args.color)
			test.AssertEqual(t, got, tt.want, "check result")
		})
	}
}
