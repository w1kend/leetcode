package permitationinstring

import (
	"fmt"
	"leetcode/go/test"
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{
			name: "e1",
			s1:   "ab",
			s2:   "eidbaooo",
			want: true,
		},
		{
			name: "e2",
			s1:   "ab",
			s2:   "eidboaoo",
			want: false,
		},
		{
			name: "e3",
			s1:   "adc",
			s2:   "dcda",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			test.AssertEqual(t, got, tt.want, fmt.Sprintf("s1-%s,s2-%s", tt.s1, tt.s2))
		})
	}
}
