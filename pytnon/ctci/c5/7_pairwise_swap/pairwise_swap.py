import unittest


def swap_pairs(num: int) -> int:
    even_ones = int("A" * 16, 16)  # "A" in hex is 1010 in binary
    odd_ones = even_ones >> 1

    even_bits_of_num = num & even_ones
    odd_bits_of_num = num & odd_ones
    return (odd_bits_of_num << 1) | (even_bits_of_num >> 1)


class Test(unittest.TestCase):
    def test_swap_pairs(self):
        tests = [
            (int('101010', 2), int('010101', 2)),
            (int('111001', 2), int('110110', 2))
        ]

        for num, want in tests:
            self.assertEqual(want, swap_pairs(num))
