package list

import (
	"reflect"
	"testing"
)

func TestFromSlice(t *testing.T) {
	type args struct {
		s []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_ToSlice(t *testing.T) {
	type fields struct {
		Val  int
		Next *ListNode
		Prev *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
				Prev: tt.fields.Prev,
			}
			if got := l.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListNode.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
