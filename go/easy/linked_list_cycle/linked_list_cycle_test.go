package linkedlistcycle

import (
	"leetcode/go/pkg/list"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *list.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
