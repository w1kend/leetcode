import unittest

from solutions import add_two_numbers
from solutions import longest_substr_wo_repeating as lswr
from solutions import longest_palindromic_substr as lps
from solutions import reverse_int32 as ri32
from solutions import string_to_int as sti
from solutions import browser_history


class TestMediumTasks(unittest.TestCase):
    # add two numbers
    def test_add_two_numbers(self):
        # l1 = [2,4,3], l2 = [5,6,4] => [7,0,8]
        l1 = add_two_numbers.ListNode().fromList([2, 4, 3])
        l2 = add_two_numbers.ListNode().fromList([5, 6, 4])
        want = add_two_numbers.ListNode().fromList([7, 0, 8])
        self.assertEqual(add_two_numbers.Solution().addTwoNumbers(
            l1, l2).list(), want.list())
        # l1 = [9,9,9,9,9,9,9] l2= [9,9,9,9] => [8,9,9,9,0,0,0,1]
        l1 = add_two_numbers.ListNode().fromList([9, 9, 9, 9, 9, 9, 9])
        l2 = add_two_numbers.ListNode().fromList([9, 9, 9, 9])
        want = add_two_numbers.ListNode().fromList([8, 9, 9, 9, 0, 0, 0, 1])
        self.assertEqual(add_two_numbers.Solution().addTwoNumbers(
            l1, l2).list(), want.list())

    def test_longest_substr_wo_repeating(self):
        self.assertEqual(
            lswr.Solution().lengthOfLongestSubstring('abcabcbb'), 3)
        self.assertEqual(lswr.Solution().lengthOfLongestSubstring('bbbbb'), 1)
        self.assertEqual(lswr.Solution().lengthOfLongestSubstring('pwwkew'), 3)
        self.assertEqual(lswr.Solution().lengthOfLongestSubstring(' '), 1)
        self.assertEqual(lswr.Solution().lengthOfLongestSubstring('au'), 2)
        self.assertEqual(
            lswr.Solution().lengthOfLongestSubstring('abcdefaghjtr'), 11)

    def test_longest_palindromic_substring(self):
        tests = {
            'babad': 'bab',
            'cbbd': 'bb',
            'abcdcbaaaabc': 'cbaaaabc',
            'abcdcbaaabcdcb': 'bcdcbaaabcdcb',
        }

        for s, want in tests.items():
            self.assertEqual(
                lps.Solution().longestPalindrome(s),
                want,
                "unexpected result for {0}".format(s)
            )

    def test_reverse_int32_integer(self):
        tests = {
            123: 321,
            -123: -321,
            120: 21,
            -8463847412: -2147483648,
            -9463847412: 0,
            2147483647: 0,
            7463847412: 2147483647,
            8463847412: 0,
        }

        for val, want in tests.items():
            self.assertEqual(
                ri32.Solution().reverse(val),
                want,
                "unexpected reversed int"
            )

    def test_string_to_int(self):
        tests = {
            'words and 987': 0,
            '': 0,
            '   -42': -42,
            '42': 42,
            '-+12': 0,
            '-2147483648': -2147483648,
            '2147483648': 2147483647,
            '    -2147483649': -2147483648,
            '   01234': 1234,
        }

        for val, want in tests.items():
            self.assertEqual(sti.Solution().myAtoi(val),
                             want, 'unexpected int')

    def test_browser_history(self):
        bh = browser_history.BrowserHistory("leetcode.com")
        bh.visit('google.com')
        bh.visit('facebook.com')
        bh.visit('youtube.com')

        self.assertEqual(bh.back(1), 'facebook.com', '')
        self.assertEqual(bh.back(1), 'google.com', '')
        self.assertEqual(bh.forward(1), 'facebook.com', '')

        bh.visit('linkedin.com')
        self.assertEqual(bh.forward(2), 'linkedin.com', '')
        self.assertEqual(bh.back(2), 'google.com', '')
        self.assertEqual(bh.back(7), 'leetcode.com', '')


if __name__ == '__main__':
    unittest.main()
