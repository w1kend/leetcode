class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        d = dict()
        for i in range(len(nums)):
            ival = nums[i]
            if ival in d:
                return [d.get(ival), i]
            d[target-ival] = i
