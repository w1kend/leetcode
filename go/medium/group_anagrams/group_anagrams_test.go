package groupanagrams

import (
	"leetcode/go/test"
	"sort"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want [][]string
	}{
		{
			name: "e1",
			in:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			want: [][]string{
				{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"},
			},
		},
		{
			name: "e2",
			in:   []string{""},
			want: [][]string{{""}},
		},
		{
			name: "e3",
			in:   []string{"a"},
			want: [][]string{{"a"}},
		},
		{
			name: "e4",
			in:   []string{"ac", "c"},
			want: [][]string{
				{"c"}, {"ac"},
			},
		},
		{
			name: "e5",
			in:   []string{"ddddddddddg", "dgggggggggg"},
			want: [][]string{
				{"ddddddddddg"}, {"dgggggggggg"},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := groupAnagrams(tt.in)

			sortMatrix(got)
			sortMatrix(tt.want)

			test.AssertEqual(t, got, tt.want, "")
		})
	}
}

func sortMatrix(s [][]string) {
	for _, sub := range s {
		sort.Slice(sub, func(i, j int) bool {
			return sub[i] < sub[j]
		})
	}

	sort.Slice(s, func(i, j int) bool {
		return s[i][0] < s[j][0]
	})
}
