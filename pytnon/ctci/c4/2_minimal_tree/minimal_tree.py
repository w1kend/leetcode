import unittest
from typing import List

from pytnon.btree import BTreeNode


def minimal_height_tree_from_array(array: List[int]) -> BTreeNode:
    root = BTreeNode(0)

    def fill(node: BTreeNode, start, end):
        if start == end:
            node.value = array[start]
            return
        mid = (start + end) // 2
        node.value = array[mid]

        if start <= mid - 1:
            node.left = BTreeNode(0)
            fill(node.left, start, mid - 1)

        if end >= mid + 1:
            node.right = BTreeNode(0)
            fill(node.right, mid + 1, end)

        return

    fill(root, 0, len(array) - 1)
    return root


class Test(unittest.TestCase):
    def test_min_heigth_tree_from_array(self):
        tests = [
            (
                [1, 2, 3, 4, 5],
                BTreeNode(
                    3,
                    BTreeNode(1, None, BTreeNode(2)),
                    BTreeNode(4, None, BTreeNode(5)),
                )
            )
        ]

        for [arr, expected] in tests:
            got = minimal_height_tree_from_array(arr)
            self.assertEqual(str(expected), str(got), f'expected - {expected}\ngot  - {got}')
