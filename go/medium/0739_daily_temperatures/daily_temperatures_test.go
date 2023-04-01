package dailytemperatures

import (
	"leetcode/go/test"
	"testing"
)

func Test_dailyTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		temps []int
		want  []int
	}{
		{
			name:  "e1",
			temps: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:  []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:  "e2",
			temps: []int{30, 40, 50, 60},
			want:  []int{1, 1, 1, 0},
		},
		{
			name:  "e3",
			temps: []int{30, 60, 90},
			want:  []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := dailyTemperatures(tt.temps)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
