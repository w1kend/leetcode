package firstbadversion

import (
	"leetcode/go/testsupport"
	"testing"
)

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n   int
		srv Srv
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 2,
				srv: Srv{
					FirstBadVersion: 2,
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 5,
				srv: Srv{
					FirstBadVersion: 4,
				},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				n: 2126753390,
				srv: Srv{
					FirstBadVersion: 1702766719,
				},
			},
			want: 1702766719,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := firstBadVersion(tt.args.n, tt.args.srv)
			testsupport.AssertEqual(t, got, tt.want, "unexpected first bad version")
		})
	}
}
