# https://leetcode.com/problems/contains-duplicate/
from ast import List


class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        return len(set(nums)) != len(nums)
