import unittest


def flip_bit(a: int) -> int:
    # check if a contains all ones
    if ~a == 0:
        return (~0).bit_length()

    # length of the previous interval of ones. the right part
    prev_length = 0
    # length of the interval of ones. the left part
    current_length = 0
    max_length = 1

    # 00111001111011111
    while a != 0:
        # last bit is 1
        if a & 1 == 1:
            current_length += 1
        else:
            # don't have the right part of the interval
            if prev_length == 0:
                max_length = max(max_length, current_length + 1)
            else:
                max_length = max(max_length, current_length + 1 + prev_length)

            prev_length = current_length
            current_length = 0

        a = a >> 1

    max_length = max(max_length, current_length + prev_length + 1)

    return max_length


class Test(unittest.TestCase):
    def test_flip_bit(self):
        tests = [
            (215, 5),  # 11010111
            (6647, 9),  # 1100111110111
            (13287, 6),  # 11001111100111
            (7143, 8),  # 1101111100111
            (2, 2),
            (1, 2),
            (3, 3),
            (~0, (~0).bit_length()),
            (0, 1),
        ]

        for value, expected in tests:
            self.assertEqual(expected, flip_bit(value), f'value - {value} ({value:016b})')
