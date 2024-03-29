package maxdifficultyjobschedule

import "testing"

func Test_minDifficulty(t *testing.T) {
	type args struct {
		jobDifficulty []int
		d             int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "e1",
			args: args{
				jobDifficulty: []int{6, 5, 4, 3, 2, 1},
				d:             2,
			},
			want: 7,
		},
		{
			name: "e2",
			args: args{
				jobDifficulty: []int{9, 9, 9},
				d:             4,
			},
			want: -1,
		},
		{
			name: "e3",
			args: args{
				jobDifficulty: []int{1, 1, 1},
				d:             3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifficulty(tt.args.jobDifficulty, tt.args.d); got != tt.want {
				t.Errorf("minDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
