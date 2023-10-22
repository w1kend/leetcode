import unittest

from pytnon.stack import Stack


class QueueStack:

    def __init__(self):
        self.__in_stack = Stack()
        self.__out_stack = Stack()

    def put(self, value):
        self.__in_stack.put(value)

    def pop(self):
        self.__fill_out_stack()
        if self.__out_stack:
            return self.__out_stack.pop()
        return None

    def peek(self):
        self.__fill_out_stack()
        if self.__out_stack:
            return self.__out_stack.peek()
        return None

    def __fill_out_stack(self):
        if not self.__out_stack:
            while self.__in_stack:
                self.__out_stack.put(self.__in_stack.pop())


class Test(unittest.TestCase):
    def test_queue_using_stacks(self):
        q = QueueStack()
        q.put(1)
        q.put(2)

        self.assertEqual(1, q.peek())

        q.put(3)
        self.assertEqual(1, q.peek())

        self.assertEqual(1, q.pop())
        self.assertEqual(2, q.pop())
        self.assertEqual(3, q.pop())
