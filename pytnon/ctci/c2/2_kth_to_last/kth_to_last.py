import unittest

from pytnon.ctci.c2 import LinkedList


def kth_to_last_element(root: LinkedList.LinkedList, k: int) -> LinkedList.LinkedListNode:
    follower: LinkedList.LinkedListNode = root.head
    dist = 2
    leader = follower.next

    if leader is None:
        return follower

    # create distance between two nodes
    while leader is not None and dist < k:
        leader = leader.next
        dist += 1

    # find the end of the list keeping K distance
    while leader.next is not None:
        leader = leader.next
        follower = follower.next

    return follower


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
