import unittest

from pytnon.ctci.c2 import LinkedList


# Write code to remove duplicates from an unsorted linked list
def remove_deps_hashmap(root: LinkedList.LinkedList) -> LinkedList.LinkedList:
    values = dict()

    curr = root.head
    values[curr.value] = True

    while curr.next:
        if curr.next.value in values:
            curr.next = curr.next.next
        else:
            values[curr.next.value] = True
            curr = curr.next

    return root


class Test(unittest.TestCase):
    def test_remove_deps(self):
        tests = [
            (LinkedList.LinkedList([1, 1, 2, 2, 2, 3, 4, 5]), LinkedList.LinkedList([1, 2, 3, 4, 5])),
            (LinkedList.LinkedList([1, 1, 1]), LinkedList.LinkedList([1])),
            (LinkedList.LinkedList([1]), LinkedList.LinkedList([1])),
        ]

        for [l, expect] in tests:
            self.assertEqual(str(expect), str(remove_deps_hashmap(l)))
