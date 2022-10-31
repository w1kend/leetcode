import unittest

from add_two_numbers import ListNode, Solution


class TestAddTwoNumbers(unittest.TestCase):
    def test_add_two_numbers(self):
        # l1 = [2,4,3], l2 = [5,6,4] => [7,0,8]
        l1 = ListNode().fromList([2, 4, 3])
        l2 = ListNode().fromList([5, 6, 4])
        want = ListNode().fromList([7, 0, 8])
        self.assertEqual(Solution().addTwoNumbers(l1, l2).list(), want.list())
        # l1 = [9,9,9,9,9,9,9] l2= [9,9,9,9] => [8,9,9,9,0,0,0,1]
        l1 = ListNode().fromList([9, 9, 9, 9, 9, 9, 9])
        l2 = ListNode().fromList([9, 9, 9, 9])
        want = ListNode().fromList([8, 9, 9, 9, 0, 0, 0, 1])
        self.assertEqual(Solution().addTwoNumbers(l1, l2).list(), want.list())


if __name__ == '__main__':
    unittest.main()
