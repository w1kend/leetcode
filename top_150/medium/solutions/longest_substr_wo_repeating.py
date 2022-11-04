from collections import defaultdict
from turtle import pos


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        positions = defaultdict(int)
        first_occur_idx: int = 0
        max_len: int = 0
        # abcabcbb
        for i, char in enumerate(s):
            # calculate the start of a window where "char" is not present
            first_occur_idx = max(first_occur_idx, positions[char])
            # check length from the start to the "i" index
            max_len = max(max_len, i - first_occur_idx + 1)
            # save the last index of "char"
            # add + 1 to avoid matching with a default int value
            positions[char] = i + 1

        return max_len
