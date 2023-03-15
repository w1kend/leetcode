import unittest

from solutions import median_two_sorted_arrays as mtsa


class TestHardTasks(unittest.TestCase):
    def test_median_two_sorted_arrays(self):
        self.assertEqual(
            mtsa.Solution().findMedianSortedArrays([1, 3], [2]), 2)
        self.assertEqual(
            mtsa.Solution().findMedianSortedArrays([1, 2], [3, 4]), 2.5)
        self.assertEqual(
            mtsa.Solution().findMedianSortedArrays([], [1]), 1)
        self.assertEqual(
            mtsa.Solution().findMedianSortedArrays([1, 2, 3], [1, 2]), 2)


if __name__ == '__main__':
    unittest.main()
