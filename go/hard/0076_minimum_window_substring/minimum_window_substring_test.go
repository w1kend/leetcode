package minimumwindowsubstring

import (
	"fmt"
	"leetcode/go/test"
	"testing"
)

func Test_minWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "e1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "e2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "e3",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "e4",
			s:    "ab",
			t:    "a",
			want: "a",
		},
		{
			name: "e5",
			s:    "ab",
			t:    "b",
			want: "b",
		},
		{
			name: "e6",
			s:    "ab",
			t:    "A",
			want: "",
		},
		{
			name: "e7",
			s:    "abc",
			t:    "b",
			want: "b",
		},
		{
			name: "m8",
			s:    "abcdef",
			t:    "bcd",
			want: "bcd",
		},
		{
			name: "e9",
			s:    "abc",
			t:    "ac",
			want: "abc",
		},
		{
			name: "e10",
			s:    "adobecodebanc",
			t:    "abcda",
			want: "adobecodeba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t)
			test.AssertEqual(t, got, tt.want, fmt.Sprintf("\ncase. s - %s, t - %s\n", tt.s, tt.t))
		})
	}
}
