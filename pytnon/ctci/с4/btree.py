from __future__ import annotations

from typing import Optional


class BTreeNode:
    def __init__(self, val: int, left: Optional[BTreeNode] = None, right: Optional[BTreeNode] = None):
        self.value: int = val
        self.left: Optional[BTreeNode] = left
        self.right: Optional[BTreeNode] = right

    def __str__(self):
        # todo: convert to list
        s = f'val={self.value}'
        if self.left:
            s += f',left=({self.left})'
        if self.right:
            s += f',right({self.right})'

        return s


class BTree:
    def __init__(self):
        self.root: Optional[BTreeNode] = None
