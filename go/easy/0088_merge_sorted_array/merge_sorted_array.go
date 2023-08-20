package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	idx := m + n + 1

	for ; idx >= 0; idx-- {
		switch {
		case m < 0:
			nums1[idx] = nums2[n]
			n--
		case n < 0:
			nums1[idx] = nums1[m]
			m--
		case nums1[m] >= nums2[n]:
			nums1[idx] = nums1[m]
			m--
		default:
			nums1[idx] = nums2[n]
			n--
		}
	}
}
