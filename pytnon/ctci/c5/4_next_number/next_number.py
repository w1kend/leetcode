import unittest


def next_number(number: int) -> (int, int):
    # change first (backwards) 0 to 1 and closest to the left 1 to 0 to get smaller value
    smaller = lower_value_with_the_same_number_of_ones(number)
    # change first (backwards) 0 to 1 and closest to the right 1 to 0 to get bigger value
    bigger = bigger_value_with_the_same_number_of_ones(number)

    return smaller, bigger


def lower_value_with_the_same_number_of_ones(num: int) -> int:
    if num == 0:
        return 0
    tmp = num
    zero_idx = -1
    one_idx = -1

    idx = 0
    while tmp != 0:
        if tmp & 1 == 0:
            zero_idx = idx

        if tmp & 1 == 1:
            one_idx = idx
            if zero_idx >= 0:
                break

        idx += 1
        tmp >>= 1

    # didn't find zero
    if zero_idx == -1:
        # there is no lower value with the same number of ones
        return num

    num = update_bit(num, zero_idx, True)
    return update_bit(num, one_idx, False)


def bigger_value_with_the_same_number_of_ones(num: int) -> int:
    if num == 0:
        return 0
    tmp = num
    last_zero_idx = -1
    closest_one_idx = -1

    idx = 0
    while tmp != 0:
        if tmp & 1 == 0:
            last_zero_idx = idx
            if closest_one_idx >= 0:
                break

        if tmp & 1 == 1:
            closest_one_idx = idx

        idx += 1
        tmp >>= 1

    # didn't find zero
    if last_zero_idx == -1:
        # add 1 before last bit. todo: check overflow
        last_zero_idx = closest_one_idx + 1

    num = update_bit(num, last_zero_idx, True)

    return update_bit(num, closest_one_idx, False)


def update_bit(num: int, idx: int, one: bool) -> int:
    num &= ~(1 << idx)  # set a bit at idx to zero
    if not one:
        return num

    return num | (1 << idx)


class Test(unittest.TestCase):
    def test_update_bit(self):
        tests = [
            (int('1011', 2), 2, True, int('1111', 2)),
            (int('1011', 2), 2, False, int('1011', 2)),
            (int('1111', 2), 2, False, int('1011', 2)),
            (int('1111', 2), 0, False, int('1110', 2)),
        ]

        for val, idx, is_one, expected in tests:
            self.assertEqual(expected, update_bit(val, idx, is_one))

    def test_biggest_mask(self):
        tests = [
            (int('1111', 2), int('10111', 2)),
            (int('1101', 2), int('1110', 2)),
            (int('000', 2), int('000', 2)),
            (int('1', 2), int('10', 2)),
        ]

        for val, expected in tests:
            self.assertEqual(expected, bigger_value_with_the_same_number_of_ones(val))

    def test_lower_value(self):
        tests = [
            (int('1111', 2), int('1111', 2)),
            (int('1011', 2), int('0111', 2)),
            (int('11011011', 2), int('11010111', 2)),
            (int('11011010', 2), int('11011001', 2)),
            (int('11011000', 2), int('11010100', 2)),
        ]

        for val, expected in tests:
            self.assertEqual(expected, lower_value_with_the_same_number_of_ones(val))
