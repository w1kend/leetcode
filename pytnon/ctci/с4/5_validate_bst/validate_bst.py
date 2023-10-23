import unittest
from typing import Optional

from pytnon.btree import BTreeNode


def validate_bst(root: BTreeNode) -> bool:
    def is_valid(node: BTreeNode, top: Optional[int], bottom: Optional[int]) -> bool:
        if node is None:
            return True

        # if current value is bigger than top limit
        if top is not None and node.value >= top:
            return False

        # if current value is less than bottom limit
        if bottom is not None and node.value <= bottom:
            return False

        # check left subtree
        if not is_valid(node.left, node.value, bottom):
            return False

        # check right subtree
        return is_valid(node.right, top, node.value)

    return is_valid(root, None, None)


class Test(unittest.TestCase):
    def test_validate_bst(self):
        tests = [
            (
                BTreeNode(8),
                True,
            ),
            (
                BTreeNode(8, BTreeNode(7), BTreeNode(10)),
                True,
            ),
            (
                BTreeNode(8, BTreeNode(2), BTreeNode(7)),
                False,
            ),
            (
                BTreeNode(
                    8,
                    BTreeNode(3, BTreeNode(1), BTreeNode(6, BTreeNode(4), BTreeNode(7))),
                    BTreeNode(10, None, BTreeNode(14, BTreeNode(13)))
                ),
                True,
            ),
            (
                BTreeNode(2, BTreeNode(2), BTreeNode(2)),
                False,
            ),
            (
                BTreeNode(0, None, BTreeNode(-1)),
                False
            ),
        ]

        for [root, expected] in tests:
            self.assertEqual(expected, validate_bst(root))
