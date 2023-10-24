import unittest

from pytnon.LinkedList import LinkedList


def intersected_node(first: LinkedList, second: LinkedList):
    for node in first:
        for sec_node in second:
            if node is sec_node:
                return node

    return None


def intersected_node_linear(first: LinkedList, second: LinkedList):
    # calc length diff
    first_len = list_len(first)
    second_len = list_len(second)

    if first_len >= second_len:
        big = first.head
        small = second.head
    else:
        big = second.head
        small = first.head

    for i in range(abs(first_len - second_len)):
        # move node in a bigger list, to make them "equal" by length
        big = big.next

    while big:
        if big is small:
            return big
        big = big.next
        small = small.next

    return None


def list_len(l: LinkedList) -> int:
    length = 0
    node = l.head
    while node:
        length += 1
        node = node.next

    return length


class Test(unittest.TestCase):
    def test_intersect_node(self):
        shared_nodes = [LinkedList([1, 2, 3]), LinkedList([2, 3])]

        tests = [
            (
                self.build_list(LinkedList([1, 2]), shared_nodes[0]),
                self.build_list(LinkedList([3, 4, 1, 2]), shared_nodes[0]),
                shared_nodes[0].head,
            ),
            (
                self.build_list(LinkedList([1, 2]), LinkedList([3, 2])),
                self.build_list(LinkedList([3, 4, 1, 2]), shared_nodes[0]),
                None,
            ),
            (
                self.build_list(LinkedList([3, 4, 1, 2]), shared_nodes[0]),
                self.build_list(LinkedList([1, 2]), shared_nodes[0]),
                shared_nodes[0].head,
            ),
        ]

        for [first, second, expected] in tests:
            self.assertIs(expected, intersected_node(first, second))
            self.assertIs(expected, intersected_node_linear(first, second))

    def build_list(self, left: LinkedList, right: LinkedList) -> LinkedList:
        left.tail.next = right.head
        left.tail = right.tail
        return left
