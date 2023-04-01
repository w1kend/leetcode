package asteroidcollision

import (
	"fmt"
	"leetcode/go/test"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	tests := []struct {
		name      string
		asteroids []int
		want      []int
	}{
		{
			name:      "e1",
			asteroids: []int{5, 10, -5},
			want:      []int{5, 10},
		},
		{
			name:      "e2",
			asteroids: []int{8, -8},
			want:      []int{},
		},
		{
			name:      "e3",
			asteroids: []int{10, 2, -5},
			want:      []int{10},
		},
		{
			name:      "e4",
			asteroids: []int{-2, -1, 1, 2},
			want:      []int{-2, -1, 1, 2},
		},
		{
			name:      "e5",
			asteroids: []int{1, -2, -2, -2},
			want:      []int{-2, -2, -2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := asteroidCollision(tt.asteroids)
			test.AssertEqual(t, got, tt.want, fmt.Sprintf("case - %v", tt.asteroids))
		})
	}
}
