import unittest


def insertion(n: int, m: int, i: int, j: int) -> int:
    # clear beats in n between i and j
    for idx in range(i, j + 1):
        n = clear_bit(n, idx)

    # move m bits to the right position
    m = m << i
    # place them in the n with OR
    return n | m


def clear_bit(val: int, i: int) -> int:
    return val & (~(1 << i))


class Test(unittest.TestCase):
    def test_insertion(self):
        tests = [
            (1024, 19, 2, 6, 1100),
            (255, 10, 3, 6, 215)
        ]

        for n, m, i, j, expected in tests:
            self.assertEqual(expected, insertion(n, m, i, j))
