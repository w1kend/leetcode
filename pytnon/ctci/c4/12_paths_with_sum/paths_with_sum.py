import time
import unittest
from typing import Optional

from pytnon.btree import BTreeNode


def paths_with_sum(root: Optional[BTreeNode], sum: int) -> int:
    if root is None:
        return 0
    return find_paths(root, 0, sum) + paths_with_sum(root.left, sum) + paths_with_sum(root.right, sum)


def find_paths(node: Optional[BTreeNode], curr_sum: int, target_sum: int) -> int:
    if node is None:
        return 0

    curr_sum += node.value
    cnt = 0
    if curr_sum == target_sum:
        cnt += 1

    cnt_left = find_paths(node.left, curr_sum, target_sum)
    cnt_right = find_paths(node.right, curr_sum, target_sum)

    return cnt_right + cnt_left + cnt


def optimized(
        node: Optional[BTreeNode],
        running_sum: Optional[int],
        target_sum: int,
        cache: dict[int, int] = None,
) -> int:
    if cache is None:
        cache = {}

    if node is None:
        return 0

    if running_sum is None:
        running_sum = node.value
    else:
        running_sum += node.value

    total_paths = 1 if running_sum == target_sum else 0
    diff = running_sum - target_sum
    total_paths += cache.get(diff, 0)

    # cache would look like {running_sum:frequency}
    if running_sum not in cache:
        cache[running_sum] = 0
    cache[running_sum] += 1

    total_paths += optimized(node.left, running_sum, target_sum, cache)
    total_paths += optimized(node.right, running_sum, target_sum, cache)

    cache[running_sum] -= 1

    return total_paths


class Test(unittest.TestCase):
    tests = [
        (
            BTreeNode(
                1,
                BTreeNode(
                    -1,
                    BTreeNode(2, BTreeNode(-1, BTreeNode(4), BTreeNode(-1))),
                    BTreeNode(1, BTreeNode(-1))
                ),
                BTreeNode(2, BTreeNode(-5, BTreeNode(2), BTreeNode(4)), BTreeNode(-3))
            ),
            0,
            9,
        )
    ]

    #
    def setUp(self):
        self.startTime = time.time()

    def tearDown(self):
        diff = time.time() - self.startTime
        print(f"{self.__module__}.{self._testMethodName}: {diff: .6f}")

    def test_paths_with_sum(self):
        for tree, target_sum, expected in self.tests:
            self.assertEqual(expected, paths_with_sum(tree, target_sum))

    def test_path_with_sum_optimized(self):
        for tree, target_sum, expected in self.tests:
            self.assertEqual(expected, optimized(tree, None, target_sum))
