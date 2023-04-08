package medianoftwosortedarrays

import (
	"leetcode/go/pkg/sugar/cmpr"
)

// the solution is to choose the right left part of each array
// if we merge these parts we have to get the left part of the whole merged and sorted array.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return median(nums2)
	}
	if len(nums2) == 0 {
		return median(nums1)
	}

	lLen := len(nums1)
	rLen := len(nums2)
	halfLen := (lLen + rLen) / 2

	r := rLen / 2
	l := halfLen - r

	for {
		if r > 0 && l != lLen && nums2[r-1] > nums1[l] {
			// increase part from left slice
			r--
			l++
			continue
		}

		if l > 0 && r != rLen && nums1[l-1] > nums2[r] {
			// increase part from right slice
			l--
			r++
			continue
		}
		break
	}

	var lMax int
	switch {
	case l == 0:
		lMax = nums2[r-1]
	case r == 0:
		lMax = nums1[l-1]
	default:
		lMax = cmpr.Max(nums1[l-1], nums2[r-1])
	}

	var rMin int
	switch {
	case r == rLen:
		rMin = nums1[l]
	case l == lLen:
		rMin = nums2[r]
	case r < rLen && l < lLen:
		rMin = cmpr.Min(nums1[l], nums2[r])
	}

	if (lLen+rLen)%2 == 0 {
		return (float64(lMax) + float64(rMin)) / 2
	}

	return float64(rMin)
}

func median(arr []int) float64 {
	mid := len(arr) / 2
	if len(arr)%2 == 0 {
		return (float64(arr[mid-1]) + float64(arr[mid])) / 2
	}

	return float64(arr[mid])
}
