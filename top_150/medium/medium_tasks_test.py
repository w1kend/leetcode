import unittest

from solutions import add_two_numbers
from solutions import longest_substr_wo_repeating as lswr


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


if __name__ == '__main__':
    unittest.main()
