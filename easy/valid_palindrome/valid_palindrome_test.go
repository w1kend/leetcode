package validpalindrome

import (
	"leetcode/testsupport"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			args: args{
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "qw@3ae",
			args: args{
				s: "qw@3ae",
			},
			want: false,
		},
		{
			name: "0P",
			args: args{
				s: "0P",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.args.s)
			testsupport.AssertEqual(t, got, tt.want, "check result")
		})
	}
}
