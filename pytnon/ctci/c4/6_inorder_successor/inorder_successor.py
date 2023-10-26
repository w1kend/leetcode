import unittest
from typing import Optional

from pytnon.btree import BTreeNode


def inorders_successor(root: BTreeNode, x: BTreeNode) -> Optional[BTreeNode]:
    result_node: Optional[BTreeNode] = None

    def in_order_traversal(node: BTreeNode):
        nonlocal result_node

        if node is None:
            return

        in_order_traversal(node.left)
        if result_node is not None:
            return

        if node.value > x.value:
            result_node = node
            return

        in_order_traversal(node.right)

    in_order_traversal(root)
    return result_node


def inorder_successor_linear(root: BTreeNode, x: BTreeNode) -> Optional[BTreeNode]:
    successor: Optional[BTreeNode] = None

    while root is not None:
        if root.value > x.value:
            successor = root
            root = root.left
            continue

        if root.value <= x.value:
            root = root.right
            continue

    return successor


class Test(unittest.TestCase):
    def test_inorder_successor(self):
        tree = BTreeNode(
            20,
            BTreeNode(
                10,
                BTreeNode(
                    5,
                    BTreeNode(1),
                    BTreeNode(7),
                ),
                BTreeNode(
                    15,
                    BTreeNode(13, BTreeNode(12), BTreeNode(14)),
                    BTreeNode(19),

                )
            ),
            BTreeNode(
                25,
                BTreeNode(23),
                BTreeNode(
                    29,
                    BTreeNode(27),
                    BTreeNode(50, BTreeNode(40), BTreeNode(60)),
                ),
            )
        )

        self.assertEqual(15, inorders_successor(tree, BTreeNode(14)).value)
        self.assertEqual(15, inorder_successor_linear(tree, BTreeNode(14)).value)
