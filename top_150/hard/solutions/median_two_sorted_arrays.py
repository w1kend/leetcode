import sys
from typing import List

# https://leetcode.com/problems/median-of-two-sorted-arrays/


class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        n = len(nums1)
        m = len(nums2)

        if n == 0:
            return nums2[int(m/2)] if m % 2 == 1 else (nums2[int(m/2-1)] + nums2[int(m/2)])/2

        if m == 0:
            return nums1[int(n/2)] if n % 2 == 1 else (nums1[int(n/2-1)] + nums1[int(n/2)])/2

        if n > m:
            return self.findMedianSortedArrays(nums2, nums1)

        start = 0
        end = n
        mid_in_merged = int((n+m)/2)

        while start < end:
            mid = int((start+end)/2)
            rmid = mid_in_merged - mid

            if nums1[mid] < nums2[rmid-1]:
                start = mid + 1
            else:
                end = mid

        first = start
        second = mid_in_merged - first

        shortLeft = nums1[first-1] if first > 0 else -sys.maxsize
        shortRight = nums1[first] if first < n else sys.maxsize

        longLeft = nums2[second-1] if second > 0 else -sys.maxsize
        longRight = nums2[second] if second < m else sys.maxsize

        if (n+m) % 2 == 0:
            return (max(shortLeft, longLeft) + min(shortRight, longRight))/2
        return min(shortRight, longRight)
