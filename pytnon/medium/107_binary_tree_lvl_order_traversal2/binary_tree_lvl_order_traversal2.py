import queue
from typing import Optional, List

from pytnon.btree import BTreeNode


def level_order_bottom(root: Optional[BTreeNode]) -> List[List[int]]:
    if root is None:
        return [[]]

    def calc_depth(node: Optional[BTreeNode]) -> int:
        if node is None:
            return 0

        left = calc_depth(node.left)
        right = calc_depth(node.right)

        return max(left, right) + 1

    depth = calc_depth(root)
    lists = [[] for _ in range(depth)]

    nodes_queue = queue.Queue[BTreeNode]()
    nodes_queue.put(root)
    level = depth  # fill lists in reverse order
    while nodes_queue.qsize() != 0:
        level_length = nodes_queue.qsize()
        lists[level] = [0 for _ in range(level_length)]  # initialize list with level length size

        for i in range(level_length):
            node = nodes_queue.get()
            lists[level][i] = node.value
            if node.left:
                nodes_queue.put(node.left)
            if node.right:
                nodes_queue.put(node.right)

        level -= 1

    return lists
