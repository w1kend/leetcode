package longestrepeatingcharacterreplacement

import (
	"leetcode/go/test"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		name string
		str  string
		k    int
		want int
	}{
		{
			name: "e1",
			str:  "ABAB",
			k:    2,
			want: 4,
		},
		{
			name: "e2",
			str:  "AABABBA",
			k:    1,
			want: 4,
		},
		{
			name: "m3",
			str:  "AABAABBAAAAAAA",
			k:    2,
			want: 11,
		},
		{
			name: "e4",
			str:  "AAAB",
			k:    0,
			want: 3,
		},
		{
			name: "e5",
			str:  "ABCDE",
			k:    1,
			want: 2,
		},
		{
			name: "e6",
			str:  "ABBB",
			k:    2,
			want: 4,
		},
		{
			name: "e7",
			str:  "KJRGKSKKOKLPADMAGODEDNEBMJMKGAPNLSAKADRLHHDRMJTMFBSIQGHENKABISHQNRFJKEPMFIPNDNEOBRJEHFLIHMDLMCIHLHQN",
			k:    5,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.str, tt.k)
			test.AssertEqual(t, got, tt.want, tt.str)
		})
	}
}
