import unittest

from pytnon.ctci.c2 import LinkedList


def kth_to_last_element(root: LinkedList.LinkedList, k: int) -> LinkedList.LinkedListNode:
    left: LinkedList.LinkedListNode = root.head
    dist = 2
    right = left.next

    if right is None:
        return left

    # increase distance between two nodes
    while right is not None and dist < k:
        right = right.next
        dist += 1

    # find the end of the list keeping K distance
    while right.next is not None:
        right = right.next
        left = left.next

    return left


class Test(unittest.TestCase):
    def test_kth_to_last_element(self):
        tests = [
            (LinkedList.LinkedList([1, 2, 3, 4]), 2, 3),
            (LinkedList.LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), 5, 6),
            (LinkedList.LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), 10, 1),
            (LinkedList.LinkedList([2]), 1, 2),
        ]

        for [root, k, expect] in tests:
            self.assertEqual(expect, kth_to_last_element(root, k).value)
