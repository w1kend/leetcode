from collections import defaultdict
from typing import List


class FindAllAnagrams:
    def find_anagrams(self, s: str, p: str) -> List[int]:
        if len(s) < len(p):
            return []

        l_idx = 0
        r_idx = len(p) - 1

        # fill p dict
        pdict = defaultdict(int)
        for letter in p:
            pdict[letter] += 1

        window_dict = defaultdict(int)
        for letter in s[:r_idx + 1]:
            window_dict[letter] += 1

        res = []
        while r_idx < len(s) - 1:
            if pdict == window_dict:
                res.append(l_idx)

            if window_dict[s[l_idx]] > 1:
                window_dict[s[l_idx]] -= 1
            else:
                del window_dict[s[l_idx]]

            window_dict[s[r_idx + 1]] += 1
            l_idx += 1
            r_idx += 1

        if pdict == window_dict:
            res.append(l_idx)

        return res
