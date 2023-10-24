import unittest
from typing import List

from pytnon.btree import BTreeNode


def all_possible_sequences(root: BTreeNode) -> List[List[int]]:
    def find(node: BTreeNode) -> List[List[int]]:
        if node.left is None and node.right is None:
            return [[node.value]]

        left_combinations = find(node.left)
        right_combinations = find(node.right)

        current_combinations: List[List[int]] = []
        for left in left_combinations:
            for right in right_combinations:
                # add combinations for each left+right variants
                current_combinations += weave(left, right, [node.value])

        return current_combinations

    return find(root)


def weave(left: List[int], right: List[int], prefix: List[int]) -> List[List[int]]:
    if len(left) == 0 or len(right) == 0:
        return [prefix + left + right]

    left_variants = weave(left[1:], right, prefix + [left[0]])
    right_variants = weave(left, right[1:], prefix + [right[0]])

    return left_variants + right_variants


class Test(unittest.TestCase):
    def test_perm(self):
        tests = [
            (
                [0],
                [2],
                [1],
                [
                    [1, 0, 2],
                    [1, 2, 0],
                ]
            ),
            (
                [1, 0, 2],
                [5, 4],
                [3],
                [
                    [3, 1, 0, 2, 5, 4],
                    [3, 1, 0, 5, 2, 4],
                    [3, 1, 0, 5, 4, 2],
                    [3, 1, 5, 0, 2, 4],
                    [3, 1, 5, 0, 4, 2],
                    [3, 1, 5, 4, 0, 2],
                    [3, 5, 1, 0, 2, 4],
                    [3, 5, 1, 0, 4, 2],
                    [3, 5, 1, 4, 0, 2],
                    [3, 5, 4, 1, 0, 2]
                ]
            ),
        ]

        for left, right, prefix, expected in tests:
            self.assertEqual(expected, weave(left, right, prefix))

    def test_all_possible_sequences(self):
        tests = [
            (
                BTreeNode(3, BTreeNode(1, BTreeNode(0), BTreeNode(2)), BTreeNode(5, BTreeNode(4), BTreeNode(6))),
                [
                    [3, 1, 0, 2, 5, 4, 6],
                    [3, 1, 0, 5, 2, 4, 6],
                    [3, 1, 0, 5, 4, 2, 6],
                    [3, 1, 0, 5, 4, 6, 2],
                    [3, 1, 5, 0, 2, 4, 6],
                    [3, 1, 5, 0, 4, 2, 6],
                    [3, 1, 5, 0, 4, 6, 2],
                    [3, 1, 5, 4, 0, 2, 6],
                    [3, 1, 5, 4, 0, 6, 2],
                    [3, 1, 5, 4, 6, 0, 2],
                    [3, 5, 1, 0, 2, 4, 6],
                    [3, 5, 1, 0, 4, 2, 6],
                    [3, 5, 1, 0, 4, 6, 2],
                    [3, 5, 1, 4, 0, 2, 6],
                    [3, 5, 1, 4, 0, 6, 2],
                    [3, 5, 1, 4, 6, 0, 2],
                    [3, 5, 4, 1, 0, 2, 6],
                    [3, 5, 4, 1, 0, 6, 2],
                    [3, 5, 4, 1, 6, 0, 2],
                    [3, 5, 4, 6, 1, 0, 2],
                    [3, 1, 0, 2, 5, 6, 4],
                    [3, 1, 0, 5, 2, 6, 4],
                    [3, 1, 0, 5, 6, 2, 4],
                    [3, 1, 0, 5, 6, 4, 2],
                    [3, 1, 5, 0, 2, 6, 4],
                    [3, 1, 5, 0, 6, 2, 4],
                    [3, 1, 5, 0, 6, 4, 2],
                    [3, 1, 5, 6, 0, 2, 4],
                    [3, 1, 5, 6, 0, 4, 2],
                    [3, 1, 5, 6, 4, 0, 2],
                    [3, 5, 1, 0, 2, 6, 4],
                    [3, 5, 1, 0, 6, 2, 4],
                    [3, 5, 1, 0, 6, 4, 2],
                    [3, 5, 1, 6, 0, 2, 4],
                    [3, 5, 1, 6, 0, 4, 2],
                    [3, 5, 1, 6, 4, 0, 2],
                    [3, 5, 6, 1, 0, 2, 4],
                    [3, 5, 6, 1, 0, 4, 2],
                    [3, 5, 6, 1, 4, 0, 2],
                    [3, 5, 6, 4, 1, 0, 2],
                    [3, 1, 2, 0, 5, 4, 6],
                    [3, 1, 2, 5, 0, 4, 6],
                    [3, 1, 2, 5, 4, 0, 6],
                    [3, 1, 2, 5, 4, 6, 0],
                    [3, 1, 5, 2, 0, 4, 6],
                    [3, 1, 5, 2, 4, 0, 6],
                    [3, 1, 5, 2, 4, 6, 0],
                    [3, 1, 5, 4, 2, 0, 6],
                    [3, 1, 5, 4, 2, 6, 0],
                    [3, 1, 5, 4, 6, 2, 0],
                    [3, 5, 1, 2, 0, 4, 6],
                    [3, 5, 1, 2, 4, 0, 6],
                    [3, 5, 1, 2, 4, 6, 0],
                    [3, 5, 1, 4, 2, 0, 6],
                    [3, 5, 1, 4, 2, 6, 0],
                    [3, 5, 1, 4, 6, 2, 0],
                    [3, 5, 4, 1, 2, 0, 6],
                    [3, 5, 4, 1, 2, 6, 0],
                    [3, 5, 4, 1, 6, 2, 0],
                    [3, 5, 4, 6, 1, 2, 0],
                    [3, 1, 2, 0, 5, 6, 4],
                    [3, 1, 2, 5, 0, 6, 4],
                    [3, 1, 2, 5, 6, 0, 4],
                    [3, 1, 2, 5, 6, 4, 0],
                    [3, 1, 5, 2, 0, 6, 4],
                    [3, 1, 5, 2, 6, 0, 4],
                    [3, 1, 5, 2, 6, 4, 0],
                    [3, 1, 5, 6, 2, 0, 4],
                    [3, 1, 5, 6, 2, 4, 0],
                    [3, 1, 5, 6, 4, 2, 0],
                    [3, 5, 1, 2, 0, 6, 4],
                    [3, 5, 1, 2, 6, 0, 4],
                    [3, 5, 1, 2, 6, 4, 0],
                    [3, 5, 1, 6, 2, 0, 4],
                    [3, 5, 1, 6, 2, 4, 0],
                    [3, 5, 1, 6, 4, 2, 0],
                    [3, 5, 6, 1, 2, 0, 4],
                    [3, 5, 6, 1, 2, 4, 0],
                    [3, 5, 6, 1, 4, 2, 0],
                    [3, 5, 6, 4, 1, 2, 0]
                ]
            ),
        ]

        for tree, expected in tests:
            self.assertEqual(expected, all_possible_sequences(tree))
