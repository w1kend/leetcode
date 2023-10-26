import unittest
from typing import Optional

from pytnon.btree import BTreeNode


def is_subtree(first: BTreeNode, second: BTreeNode) -> bool:
    if first is None:
        return False

    if first.value == second.value and is_identical(first, second):
        return True

    return is_subtree(first.left, second.left) or is_subtree(first.right, second.right)


def is_identical(first: Optional[BTreeNode], second: Optional[BTreeNode]) -> bool:
    if first is None and second is None:
        return True

    # if some of the nodes is None or their value is different
    if (first is None or second is None) or (first.value != second.value):
        return False

    return is_identical(first.left, second.left) and is_identical(first.right, second.right)


class Test(unittest.TestCase):
    def test_is_identical(self):
        tests = [
            (
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                True,
            ),
            (
                BTreeNode(1, BTreeNode(2), BTreeNode(3, None, BTreeNode(4))),
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                False,
            ),
            (
                BTreeNode(1),
                BTreeNode(2),
                False,
            ),
            (
                BTreeNode(1),
                BTreeNode(1),
                False,
            ),
        ]

    def test_is_subtree(self):
        tests = [
            (
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                True,
            ),
            (
                BTreeNode(1, BTreeNode(2), BTreeNode(3, None, BTreeNode(4))),
                BTreeNode(1, BTreeNode(2), BTreeNode(3)),
                False,
            ),
            (
                BTreeNode(1),
                BTreeNode(2),
                False,
            ),
            (
                BTreeNode(1),
                BTreeNode(1),
                False,
            ),
            (
                BTreeNode(
                    1,
                    BTreeNode(
                        2,
                        BTreeNode(4, BTreeNode(10, BTreeNode(11), BTreeNode(12))),
                        BTreeNode(5, BTreeNode(13))
                    ),
                    BTreeNode(3, BTreeNode(6, BTreeNode(8), BTreeNode(9)), BTreeNode(7))
                ),
                BTreeNode(3, BTreeNode(6, BTreeNode(8), BTreeNode(9)), BTreeNode(7)),
                True,
            ),
            (
                BTreeNode(
                    1,
                    BTreeNode(
                        2,
                        BTreeNode(4, BTreeNode(10, BTreeNode(11), BTreeNode(12))),
                        BTreeNode(5, BTreeNode(13))
                    ),
                    BTreeNode(3, BTreeNode(6, BTreeNode(8), BTreeNode(9)), BTreeNode(7))
                ),
                BTreeNode(3, BTreeNode(6, BTreeNode(8), BTreeNode(9))),
                False,
            ),
        ]
