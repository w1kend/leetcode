package longestpalindromicsubstring

import "testing"

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "e1",
			str:  "babed",
			want: "bab",
		},
		{
			name: "e1",
			str:  "cbbd",
			want: "bb",
		},
		{
			name: "m3",
			str:  "qwerwetabcdcbacvbcvb",
			want: "abcdcba",
		},
		{
			name: "m4",
			str:  "nzjklxcjaskldkasdhousesuohdshgjkkdallqwecoolercaracrelooclaklfjdjsnnmcxzv",
			want: "coolercaracrelooc",
		},
		{
			name: "e5",
			str:  "a",
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.str); got != tt.want {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}

			if got := dpSolution(tt.str); got != tt.want {
				t.Errorf("dpSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
