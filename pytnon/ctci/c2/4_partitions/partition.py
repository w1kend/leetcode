import unittest

from pytnon.ctci.c2.LinkedList import LinkedList


# Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes
# greater than or equal to x. If x is contained within the list, the values of x only need to be after the elements
# less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to
# appear between the left and right partitions.
def partition_new_lists(root: LinkedList, n: int) -> LinkedList:
    lowerList = LinkedList()
    greaterList = LinkedList()

    curr = root.head
    while curr is not None:
        if curr.value < n:
            lowerList.add(curr.value)
        elif curr.value == n:
            greaterList.add_to_beginning(curr.value)
        else:
            greaterList.add(curr.value)

        curr = curr.next

    lowerList.tail.next = greaterList.head

    return lowerList


# move only nodes with values less than N
def partition_move_nodes(root: LinkedList, n: int) -> LinkedList:
    # at least two elements
    if root.head.next == None:
        return root

    curr = root.head
    is_bigger_partition = False

    if curr.value >= n:
        # if the value is less than N, then we just keep it in place
        # otherwise curr node is start of partitioning, and we move all <N nodes before it
        is_bigger_partition = True

    while curr.next:
        if curr.next.value < n and is_bigger_partition:
            # move it in front of the list
            tmp = curr.next
            curr.next = curr.next.next
            tmp.next = root.head
            root.head = tmp
            continue

        if curr.next.value >= n:
            is_bigger_partition = True

        curr = curr.next

    return root


class Test(unittest.TestCase):
    def test_partition(self):
        tests = [
            (LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), 5, LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])),
            (LinkedList([8, 9, 1, 2, 4, 5, 5, 10, 4, 2, 5, 11]), 5, LinkedList([1, 2, 4, 4, 2, 5, 5, 5, 8, 9, 10, 11])),
            (LinkedList([1, 4, 3]), 3, LinkedList([1, 3, 4])),
            (LinkedList([1, 2, 3]), 3, LinkedList([1, 2, 3])),
        ]

        for [root, n, expected] in tests:
            self.assertEqual(str(expected), str(partition_new_lists(root, n)))

    def test_partition_move_nodes(self):
        tests = [
            (LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]), 5, LinkedList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10])),
            (LinkedList([8, 9, 1, 2, 4, 5, 5, 10, 4, 2, 5, 11]), 5, LinkedList([2, 4, 4, 2, 1, 8, 9, 5, 5, 10, 5, 11])),
            (LinkedList([1, 4, 3, 2]), 3, LinkedList([2, 1, 4, 3])),
            (LinkedList([1, 2, 3]), 3, LinkedList([1, 2, 3])),
        ]

        for [root, n, expected] in tests:
            self.assertEqual(str(expected), str(partition_move_nodes(root, n)))
