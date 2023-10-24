import unittest

from pytnon.LinkedList import LinkedList


def sum_lists(l: LinkedList, r: LinkedList) -> LinkedList:
    l_node = l.head
    r_node = r.head

    res = LinkedList()

    overflow = False
    while l_node or r_node:
        sum_val = 0
        if l_node:
            sum_val = l_node.value
            l_node = l_node.next

        if r_node:
            sum_val += r_node.value
            r_node = r_node.next

        sum_val, overflow = check_overflow(sum_val, overflow)
        res.add(sum_val)

    if overflow:
        res.add(1)

    return res


def sum_lists_recursion(l: LinkedList, r: LinkedList) -> LinkedList:
    def sum_rec(l_node, r_node, overflow: bool, result: LinkedList) -> LinkedList:
        if l_node is None and r_node is None:
            return result
        if l_node is None:
            sum_val, overflow = check_overflow(r_node.value, overflow)
            result.add(sum_val)
            return sum_rec(None, r_node.next, overflow, result)
        if r_node is None:
            sum_val, overflow = check_overflow(l_node.value, overflow)
            result.add(sum_val)
            return sum_rec(l_node.next, None, overflow, result)

        sum_val, overflow = check_overflow(l_node.value + r_node.value, overflow)
        result.add(sum_val)
        return sum_rec(l_node.next, r_node.next, overflow, result)

    return sum_rec(l.head, r.head, False, LinkedList())


def check_overflow(val: int, overflow: bool) -> (int, bool):
    if overflow:
        val += 1

    if val >= 10:
        val -= 10
        overflow = True
    else:
        overflow = False

    return val, overflow


class Test(unittest.TestCase):
    def test_sum_lists(self):
        tests = [
            (LinkedList([1, 2, 3].reverse()), LinkedList([1, 2, 3].reverse()), LinkedList([2, 4, 6].reverse())),
            (LinkedList([6, 1.7].reverse()), LinkedList([2, 9, 5].reverse()), LinkedList([9, 1, 2].reverse())),
            (LinkedList([1].reverse()), LinkedList([1, 0, 9].reverse()), LinkedList([1, 1, 0].reverse())),
        ]

        for [first_val, second_val, expected] in tests:
            self.assertEqual(str(expected), str(sum_lists(first_val, second_val)))
            self.assertEqual(str(expected), str(sum_lists_recursion(first_val, second_val)))
