import queue
import unittest
from typing import List

from pytnon.btree import BTreeNode


def list_of_depths(root: BTreeNode) -> List[List[int]]:
    lists = []
    level = 0

    next_nodes: queue.Queue[BTreeNode] = queue.Queue()
    next_nodes.put(root)

    while next_nodes.qsize() != 0:
        # Iterating level by level.
        # During iteration, we add nodes from deeper level, so we have to keep length of current level
        level_length = next_nodes.qsize()

        # Initialize list of level_length size
        lists.append([0 for _ in range(level_length)])
        for n in range(level_length):
            node = next_nodes.get()
            lists[level][n] = node.value

            # Add nodes from deeper level to the queue
            if node.left:
                next_nodes.put(node.left)

            if node.right:
                next_nodes.put(node.right)

        level += 1

    return lists


class Test(unittest.TestCase):
    def test_list_of_depths(self):
        tests = [
            (
                BTreeNode(
                    1,
                    BTreeNode(2, BTreeNode(4), BTreeNode(5)),
                    BTreeNode(3, None, BTreeNode(7))
                ),
                [
                    [1],
                    [2, 3],
                    [4, 5, 7]
                ]
            ),
            (
                BTreeNode(
                    1,
                    BTreeNode(2, BTreeNode(4, BTreeNode(8))),
                    BTreeNode(3, BTreeNode(6), BTreeNode(7, BTreeNode(14), BTreeNode(15))),
                ),
                [
                    [1],
                    [2, 3],
                    [4, 6, 7],
                    [8, 14, 15],
                ]
            ),
            (
                BTreeNode(1),
                [[1]]
            ),
        ]

        for [root, expected] in tests:
            self.assertEqual(expected, list_of_depths(root))
