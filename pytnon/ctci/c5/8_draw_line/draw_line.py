import unittest


def draw_line(screen: bytearray, width: int, x1: int, x2: int, y) -> bytes:
    bitsize = 8

    bytes_per_line = width // bitsize

    # indexes
    start = bytes_per_line * y
    end = start + bytes_per_line - 1

    screen_line = screen[start:end + 1]

    curr_byte_num = 0  # N byte in the line
    while curr_byte_num < len(screen_line):
        byte_start = curr_byte_num * bitsize
        byte_end = (curr_byte_num + 1) * bitsize - 1
        if between(byte_start, x1, byte_end):
            # x1 starts in the current byte
            screen[start + curr_byte_num] = fill_line(screen_line[curr_byte_num], bitsize - x1,
                                                      max(1, byte_end - x2))  # reversed indexes

        if byte_start > x1:
            screen[start + curr_byte_num] = fill_line(screen_line[curr_byte_num], bitsize, max(1, byte_end - x2))

        curr_byte_num += 1

    return screen


def between(x1: int, value: int, x2: int) -> bool:
    return x1 <= value <= x2


def fill_line(byte: int, start: int, end: int) -> int:
    mask = ((1 << start) - 1) ^ ((1 << end - 1) - 1)

    return byte | mask


class Test(unittest.TestCase):
    def test_draw_line(self):
        tests = [
            (
                bytearray([int('11000000', 2), int('01100000', 2), int('01100000', 2)]),
                8,
                0, 7,
                1,
                bytearray([int('11000000', 2), int('11111111', 2), int('01100000', 2)]),
            ),
            (
                bytearray([int('11000000', 2), int('01100000', 2), int('01100000', 2)]),
                8,
                4, 6,
                0,
                bytearray([int('11001110', 2), int('01100000', 2), int('01100000', 2)]),
            ),
            (
                bytearray([int('11000000', 2), int('01100000', 2), int('01100000', 2), int('00000000', 2)]),
                16,
                4, 11,
                1,
                bytearray([int('11000000', 2), int('01100000', 2), int('01101111', 2), int('11111000', 2)]),
            ),
        ]

        for screen, width, x1, x2, y, expected in tests:
            scr = screen.copy()
            got = draw_line(screen, width, x1, x2, y)
            self.assertEqual(
                expected.hex(),
                got.hex(),
                f'screen   - {int(scr.hex(), 16): b}\nexpected - {int(expected.hex(), 16): b}\ngot      - {int(got.hex(), 16): b}',
            )
