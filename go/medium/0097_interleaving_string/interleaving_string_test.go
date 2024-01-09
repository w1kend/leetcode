package interleavingstring

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbcbcac",
			},
			want: true,
		},
		{
			name: "e2",
			args: args{
				s1: "aabcc",
				s2: "dbbca",
				s3: "aadbbbaccc",
			},
			want: false,
		},
		{
			name: "e3",
			args: args{
				s1: "a",
				s2: "b",
				s3: "a",
			},
			want: false,
		},
		{
			name: "e4",
			args: args{
				s1: "a",
				s2: "",
				s3: "a",
			},
			want: true,
		},
		{
			name: "e5",
			args: args{
				s1: "",
				s2: "",
				s3: "a",
			},
			want: false,
		},
		{
			name: "e6",
			args: args{
				s1: "",
				s2: "a",
				s3: "a",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
