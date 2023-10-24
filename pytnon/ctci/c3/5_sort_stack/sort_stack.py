import unittest

from pytnon.stack import Stack


def sort_stack(main: Stack) -> Stack:
    tmp = Stack()
    source_length = len(main)
    n = 0  # number of correctly placed elements
    while n < source_length - 1:
        if not tmp:
            tmp.put(main.pop())
        # swap them. moving the largest element to the bottom of the stack
        if main.peek() <= tmp.peek():
            tmp_val = tmp.pop()
            tmp.put(main.pop())
            main.put(tmp_val)
            if len(main) == n + 1:
                n += 1
            else:
                tmp.put(main.pop())
        elif len(main) > n:
            tmp.put(main.pop())

        if len(main) == n:
            # place all the elements back to the original stack and do again
            while tmp:
                main.put(tmp.pop())

    while tmp:
        main.put(tmp.pop())

    return main


class Test(unittest.TestCase):
    def test_sort_stack(self):
        tests = [
            (Stack([1, 4, 2, 3]), Stack([4, 3, 2, 1])),
            (Stack([1, 2, 3, 4, 5, 6]), Stack([6, 5, 4, 3, 2, 1])),
            (Stack([1, 6, 6, 5, 3, 1, 2]), Stack([6, 6, 5, 3, 2, 1, 1])),
            (Stack([1, 4, 2]), Stack([4, 2, 1])),
            (Stack([1, 2]), Stack([2, 1])),
            (Stack([2]), Stack([2])),
        ]

        for [test, expect] in tests:
            self.assertEqual(str(expect), str(sort_stack(test)))
