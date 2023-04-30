package coinchange

import (
	"fmt"
	"leetcode/go/test"
	"testing"
)

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "e1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3, // 5 + 5 + 1
		},
		{
			name:   "e2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "e3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "m4",
			coins:  []int{4, 5, 8, 10},
			amount: 11,
			want:   -1,
		},
		{
			name:   "e5",
			coins:  []int{1},
			amount: 2,
			want:   2,
		},
		{
			name:   "e6",
			coins:  []int{1, 2147483647},
			amount: 2,
			want:   2,
		},
		{
			name:   "e7",
			coins:  []int{186, 419, 83, 408},
			amount: 6249,
			want:   20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"bruteforce", func(t *testing.T) {
			got := newBrutForceChanger(tt.coins).change(tt.amount)
			test.AssertEqual(t, got, tt.want, "")
		})
		t.Run(tt.name+"iterative", func(t *testing.T) {
			c := newIterativeChanger(tt.coins)
			got := c.change(tt.amount)
			test.AssertEqual(t, got, tt.want, fmt.Sprintln(tt.coins, tt.amount, c.dp))
		})
	}
}
