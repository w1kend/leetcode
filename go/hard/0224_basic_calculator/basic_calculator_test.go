package basiccalculator

import (
	"leetcode/go/test"
	"testing"
)

func Test_calculate(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		want       int
	}{
		{
			name:       "e1",
			expression: "1 + 1",
			want:       2,
		},
		{
			name:       "e2",
			expression: "2-1 + 2 ",
			want:       3,
		},
		{
			name:       "e3",
			expression: "(1+(4+5+2)-3)+(6+8)",
			want:       23,
		},
		{
			name:       "m4",
			expression: "10+5",
			want:       15,
		},
		{
			name:       "e5",
			expression: "1-(     -2)",
			want:       3,
		},
		{
			name:       "e6",
			expression: "-(-2)+4",
			want:       6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculate(tt.expression)
			test.AssertEqual(t, got, tt.want, tt.expression)
		})
	}
}
