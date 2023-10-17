import unittest
from typing import List

from pytnon.ctci.с4.btree import BTreeNode


def minimal_height_tree_from_array(array: List[int]) -> BTreeNode:
    root = BTreeNode(0)

    def fill(node: BTreeNode, l_idx, r_idx):
        if l_idx == r_idx:
            node.value = array[l_idx]
            return
        mid = (l_idx + r_idx) // 2
        node.value = array[mid]

        if l_idx <= mid - 1:
            node.left = BTreeNode(0)
            fill(node.left, l_idx, mid - 1)

        if r_idx >= mid + 1:
            node.right = BTreeNode(0)
            fill(node.right, mid + 1, r_idx)

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
