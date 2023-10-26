import unittest


def flips_to_convert(a: int, b: int) -> int:
    diff = a ^ b

    cnt = 0
    while diff != 0:
        cnt += 1
        diff = diff & (diff - 1)  # clear the least significant bit
    return cnt


class Test(unittest.TestCase):
    def test_flips_co_convert(self):
        tests = [
            (int('11101', 2), int('01111', 2), 2),
            (int('101', 2), int('1010', 2), 4),
            (int('101', 2), int('1011', 2), 3),
            (int('101', 2), int('101', 2), 0),
        ]

        for a, b, expected in tests:
            self.assertEqual(expected, flips_to_convert(a, b))
