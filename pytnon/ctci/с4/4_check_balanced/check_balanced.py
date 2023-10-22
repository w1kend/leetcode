from typing import Optional

from pytnon.btree import BTreeNode


def check_balanced(root: Optional[BTreeNode]) -> bool:
    def tree_depth(node: Optional[BTreeNode]) -> int:
        if node is None:
            return 0

        left_depth = tree_depth(node.left)
        right_depth = tree_depth(node.right)

        if left_depth == -1 or right_depth == -1 or abs(left_depth - right_depth) > 1:
            return -1

        return max(left_depth, right_depth) + 1

    return tree_depth(root) != -1
