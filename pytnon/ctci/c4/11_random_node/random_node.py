from __future__ import annotations

import random
from collections import defaultdict
from typing import Optional


class Node:
    def __init__(self, val: int, left: Optional[Node] = None, right: Optional[Node] = None):
        self.val: int = val
        self.left: Optional[Node] = left
        self.right: Optional[Node] = right


class BinaryTree:
    def __init__(self, root: Optional[Node] = None):
        self.root = root
        self.__calc_params()

    def __calc_params(self):
        self.depth: int = self.__depth(self.root)
        self.nodes_cnt: int = self.__calc_nodes(self.root)

    def __depth(self, node: Optional[Node]) -> int:
        if node is None:
            return 0

        left_depth = self.__depth(node.left)
        right_depth = self.__depth(node.right)

        return max(left_depth, right_depth) + 1

    def __calc_nodes(self, node: Optional[Node]):
        if node is None:
            return 0

        return self.__calc_nodes(node.left) + self.__calc_nodes(node.right) + 1

    def get_random_node(self) -> Optional[Node]:
        node_num: int = random.randint(0, self.nodes_cnt - 1)

        # in-order traverse through the tree
        return self.__traverse_till(self.root, node_num)[1]

    def __traverse_till(self, node: Optional[Node], idx: int, cnt: int = 0) -> tuple[int, Optional[Node]]:
        if node is None:
            return cnt - 1, None

        if idx == cnt:
            return cnt, node

        cnt, result = self.__traverse_till(node.left, idx, cnt + 1)
        if result is not None:
            return cnt, result

        return self.__traverse_till(node.right, idx, cnt + 1)


def example():
    tree = BinaryTree(
        Node(
            1,
            Node(
                2,
                Node(4, Node(10, Node(11), Node(12))),
                Node(5, Node(13))
            ),
            Node(3, Node(6, Node(8), Node(9)), Node(7))
        ),
    )

    print(f'tree depth - {tree.depth}, nodes_cnt - {tree.nodes_cnt}')

    n = 100000
    cnts = defaultdict(int)
    for i in range(n):
        cnts[tree.get_random_node().val] += 1

    print(f'run times [{n}]\ncounts: {cnts}')


if __name__ == "__main__":
    example()
