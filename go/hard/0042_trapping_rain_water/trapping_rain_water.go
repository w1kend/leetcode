package trappingrainwater

import "leetcode/go/pkg/sugar/cmpr"

func trap(height []int) int {
	l, r := 0, len(height)-1
	amount := 0

	leftMax, rightMax := height[l], height[r]

	for l < r {
		lh, rh := height[l], height[r]
		if lh > rh {
			r--
			rightMax = cmpr.Max(rightMax, height[r])
			amount += rightMax - height[r]
		} else {
			l++
			leftMax = cmpr.Max(leftMax, height[l])
			amount += leftMax - height[l]
		}
	}

	return amount
}
