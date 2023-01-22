from typing import List


# https://leetcode.com/problems/longest-common-prefix/
class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        pref = strs[0]
        for s in strs:
            while len(pref) > 0:
                if not s.startswith(pref):
                    pref = pref[:len(pref)-1]
                else:
                    break

        return pref
