import unittest

from solutions import two_sum
from solutions import longest_common_prefix


class TestTwoSum(unittest.TestCase):
    def test_two_sum(self):
        self.assertEqual(two_sum.Solution().twoSum([2, 7, 11, 15], 9), [0, 1])
        self.assertEqual(two_sum.Solution().twoSum([3, 2, 4], 6), [1, 2])
        self.assertEqual(two_sum.Solution().twoSum([3, 3], 6), [0, 1])

    def test_longest_common_prefix(self):
        tests = [
            ['flower', 'flow', 'flight', 'fl'],
            ['dog', 'racecar', 'car', ''],
            ['c', 'acc', 'ccc', ''],
        ]

        lsp = longest_common_prefix.Solution()

        for tt in tests:
            self.assertEqual(lsp.longestCommonPrefix(
                tt[:len(tt)-1]), tt[len(tt)-1], '')


if __name__ == '__main__':
    unittest.main()
