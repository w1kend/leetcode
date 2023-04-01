package carfleet

import (
	"leetcode/go/test"
	"testing"
)

func Test_carFleet(t *testing.T) {
	type args struct {
		target   int
		position []int
		speed    []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				target:   12,
				position: []int{10, 8, 0, 5, 3}, // 0, 3, 5, 8, 10
				speed:    []int{2, 4, 1, 1, 3},  //  1, 3, 1, 4, 2
			},
			want: 3,
		},
		{
			name: "e2",
			args: args{
				target:   10,
				position: []int{3},
				speed:    []int{3},
			},
			want: 1,
		},
		{
			name: "e3",
			args: args{
				target:   100,
				position: []int{0, 2, 4},
				speed:    []int{4, 2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := carFleet(tt.args.target, tt.args.position, tt.args.speed)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
