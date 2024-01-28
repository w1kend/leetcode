package confusingnumber

import "testing"

func Test_confusingNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "e1",
			args: args{
				n: 6, // => 9
			},
			want: true,
		},
		{
			name: "m2",
			args: args{
				n: 808, // => 808
			},
			want: false,
		},
		{
			name: "",
			args: args{
				n: 89, // => 68
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := confusingNumber(tt.args.n); got != tt.want {
				t.Errorf("confusingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
