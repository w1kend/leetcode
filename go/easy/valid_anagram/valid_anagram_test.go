package validanagram

import (
	"leetcode/go/test"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "anagram - nagaram",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "rat - car",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "ab - a",
			args: args{
				s: "ab",
				t: "a",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.args.s, tt.args.t)
			test.AssertEqual(t, got, tt.want, "check results")
		})
	}
}
