package oneeditdistance

import "testing"

func Test_isOneEditDistance(t *testing.T) {
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
			name: "e1",
			args: args{
				s: "ac",
				t: "abc",
			},
			want: true,
		},
		{
			name: "e2",
			args: args{
				s: "",
				t: "",
			},
			want: false,
		},
		{
			name: "m3",
			args: args{
				s: "abc",
				t: "ab",
			},
			want: true,
		},
		{
			name: "m4",
			args: args{
				s: "abc",
				t: "avc",
			},
			want: true,
		},
		{
			name: "m5",
			args: args{
				s: "abc",
				t: "abcd",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isOneEditDistance(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isOneEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
