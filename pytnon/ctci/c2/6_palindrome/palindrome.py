import unittest

from pytnon.ctci.c2.LinkedList import LinkedList, LinkedListNode


# Implement a function to check if a linked list is a palindrome
def is_palindrome(root: LinkedList) -> bool:
    reverted_list = revert_list(root)
    curr = reverted_list.head
    for node in root:
        if curr.value != node.value:
            return False

        curr = curr.next

    return True


def revert_list(root: LinkedList) -> LinkedList:
    reverted = LinkedList()

    for node in root:
        reverted.add_to_beginning(node.value)

    return reverted


def is_palindrome_recursion(root: LinkedList, length: int) -> bool:
    def is_pal(node: LinkedListNode, level: int) -> (LinkedListNode, bool):
        if level == 0:  # even list length
            return node, True
        if level == 1:  # odd list length
            return node.next, True

        # go down to the middle element
        right_node, is_eq = is_pal(node.next, level - 2)
        if right_node is None:
            return None, is_eq

        # compare nodes on the way back. Nodes will be the same distance from the middle element
        is_eq = is_eq and right_node.value == node.value
        return right_node.next, is_eq

    return is_pal(root.head, length)[1]


class Test(unittest.TestCase):
    def test_is_palindrome(self):
        tests = [
            (LinkedList([1, 2, 3, 4, 3, 2, 1]), True),
            (LinkedList([1]), True),
            (LinkedList([1, 2]), False),
            (LinkedList([1, 2, 3, 4, 5]), False),
        ]

        for [root, expected] in tests:
            self.assertEqual(expected, is_palindrome(root))
            self.assertEqual(expected, is_palindrome_recursion(root, len(root)))
