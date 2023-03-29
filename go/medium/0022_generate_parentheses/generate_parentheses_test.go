package generateparentheses

import (
	"leetcode/go/test"
	"sort"
	"testing"
)

func Test_generateParenthesis(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "e1",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "e2",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "m3",
			n:    2,
			want: []string{"(())", "()()"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateParenthesis(tt.n)
			sort.Strings(got)
			sort.Strings(tt.want)
			test.AssertEqual(t, got, tt.want, "")
		})
	}
}
