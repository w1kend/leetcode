package containerwithmostwater

import (
	"fmt"
	"leetcode/go/pkg/sugar/cmpr"
)

func maxArea(height []int) int {
	max := 0
	xi := 0
	xj := len(height) - 1
	for xi < xj {
		yi := height[xi]
		yj := height[xj]

		area := calculateRectangleArea(xi+1, yi, xj+1, yj)
		fmt.Printf("points: (%d,%d) - (%d,%d), area - %d\n", xi+1, yi, xj+1, yj, area)
		max = cmpr.Max(max, area)

		if yi <= yj {
			xi++
		} else {
			xj--
		}
	}

	return max
}

func calculateRectangleArea(x1, y1, x2, y2 int) int {
	return (x2 - x1) * cmpr.Min(y1, y2)
}
