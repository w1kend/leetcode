package poweroftwo

import "testing"

func Test_powerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "2-true",
			n:    2,
			want: true,
		},
		{
			name: "3-false",
			n:    3,
			want: false,
		},
		{
			name: "0-false",
			n:    0,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPowerOfTwo(tt.n)
			if got != tt.want {
				t.Errorf("powerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
