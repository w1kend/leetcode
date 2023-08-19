package validpalindrome2

import (
	"leetcode/go/test"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "e1",
			s:    "aba",
			want: true,
		},
		{
			name: "e2",
			s:    "abca",
			want: true,
		},
		{
			name: "e3",
			s:    "abc",
			want: false,
		},
		{
			name: "m4",
			s:    "aaa",
			want: true,
		},
		{
			name: "m5",
			s:    "aaaxfffaaa",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validPalindrome(tt.s)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
