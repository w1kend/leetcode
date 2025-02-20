package subsets

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "123",
			nums: []int{1, 2, 3},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name: "12",
			nums: []int{1, 2},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "0",
			nums: []int{0},
			want: [][]int{{}, {0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subsets(tt.nums)
			gotm := make(map[string]struct{}, len(got))
			for _, set := range got {
				gotm[fmt.Sprintf("%v", set)] = struct{}{}
			}

			wantm := make(map[string]struct{}, len(tt.want))
			for _, set := range tt.want {
				wantm[fmt.Sprintf("%v", set)] = struct{}{}
			}

			if !reflect.DeepEqual(wantm, gotm) {
				t.Fatalf("want != got.\nwant-%v\ngot-%v", wantm, gotm)
			}
		})
	}
}
