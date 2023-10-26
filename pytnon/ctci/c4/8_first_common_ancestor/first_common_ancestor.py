import unittest
from typing import Optional

from pytnon.btree import BTreeNode


def first_common_ancestor(root: BTreeNode, p: BTreeNode, q: BTreeNode) -> BTreeNode:
    parent: Optional[BTreeNode] = None

    if p.value == q.value:
        return p

    def search(node: Optional[BTreeNode]) -> int:
        nonlocal parent

        if node is None or parent is not None:
            return 0

        curr_cnt = 0
        if node.value == p.value or node.value == q.value:
            curr_cnt = 1

        left_cnt = search(node.left)
        right_cnt = search(node.right)

        total_cnt = curr_cnt + left_cnt + right_cnt

        if parent is None and total_cnt == 2:
            parent = node
            return 0

        return total_cnt

    search(root)
    return parent


class Test(unittest.TestCase):
    def test_first_common_ancestor(self):
        """
        tree
                3
              /   \
             5     1
           /  \    | \
          6    2   0  8
              / \
             7   4
        """
        tree = BTreeNode(
            3,
            BTreeNode(5, BTreeNode(6), BTreeNode(2, BTreeNode(7), BTreeNode(4))),
            BTreeNode(1, BTreeNode(0), BTreeNode(8))
        )
        tests = [
            (
                tree,
                BTreeNode(5),
                BTreeNode(1),
                BTreeNode(3)
            ),
            (
                tree,
                BTreeNode(3),
                BTreeNode(3),
                BTreeNode(3)
            ),
            (
                tree,
                BTreeNode(6),
                BTreeNode(4),
                BTreeNode(5)
            ),
            (
                tree,
                BTreeNode(5),
                BTreeNode(6),
                BTreeNode(5)
            ),
            (
                tree,
                BTreeNode(6),
                BTreeNode(0),
                BTreeNode(3)
            ),
            (
                tree,
                BTreeNode(8),
                BTreeNode(7),
                BTreeNode(3)
            )
        ]

        for root, p, q, want in tests:
            self.assertEqual(want.value, first_common_ancestor(root, p, q).value)
