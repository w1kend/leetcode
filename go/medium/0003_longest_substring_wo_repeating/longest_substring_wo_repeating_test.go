package longestsubstringworepeating

import (
	"leetcode/go/test"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "e1",
			in:   "abcabcbb",
			want: 3,
		},
		{
			name: "e2",
			in:   "pwwkew",
			want: 3,
		},
		{
			name: "e3",
			in:   "",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.in)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
