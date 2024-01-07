package uniquepaths2

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name         string
		obstacleGrid [][]int
		want         int
	}{
		{
			name:         "e1",
			obstacleGrid: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
			want:         2,
		},
		{
			name:         "e2",
			obstacleGrid: [][]int{{0, 1}, {0, 0}},
			want:         1,
		},
		{
			name: "m3",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 0, 0},
			},
			want: 3,
		},
		{
			name: "m4",
			obstacleGrid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{1, 0, 0},
			},
			want: 5,
		},
		{
			name:         "e5",
			obstacleGrid: [][]int{{0, 0}},
			want:         1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
