package kokoeatingbananas

import (
	"leetcode/go/test"
	"testing"
)

func Test_calcHours(t *testing.T) {
	tests := []struct {
		name    string
		piles   []int
		perHour int
		want    int
	}{
		{
			name:    "m1",
			piles:   []int{8, 9, 10, 11, 12},
			perHour: 10,
			want:    7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calcHours(tt.piles, tt.perHour)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{
			name:  "e1",
			piles: []int{3, 6, 7, 11},
			h:     8,
			want:  4,
		},
		{
			name:  "e2",
			piles: []int{30, 11, 23, 4, 20},
			h:     5,
			want:  30,
		},
		{
			name:  "e3",
			piles: []int{30, 11, 23, 4, 20},
			h:     6,
			want:  23,
		},
		{
			name:  "e4",
			piles: []int{312884470},
			h:     968709470,
			want:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
