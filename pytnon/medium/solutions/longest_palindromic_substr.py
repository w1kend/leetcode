class Solution:
    source_str: str
    reversed_str: str

    def longestPalindrome(self, s: str) -> str:
        self.source_str = s
        self.reversed_str = s[::-1]

        i = 0
        j = len(s)-1

        longest = 0
        long_i = 0
        long_j = 0

        while j - i > longest:
            p = False
            while p != True and j - i > longest:
                if self.isPalindrome(i, j):
                    p = True
                else:
                    j -= 1

            if p:
                longest = j-i
                long_i = i
                long_j = j

            i += 1
            j = len(s)-1

        return s[long_i:long_j+1]

    def isPalindrome(self, i: int, j: int) -> bool:
        l = len(self.reversed_str)
        return self.source_str[i:j+1] == self.reversed_str[l-j-1: l-i]
