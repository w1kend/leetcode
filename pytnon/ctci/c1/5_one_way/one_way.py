# There are three types of edits that can be performed on strings: insert a character, remove a character,
# or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away
# pale, ple -> true; pales, pale -> true; pale, bale -> true; pale, bake -> false;
import unittest


def one_way(a: str, b: str) -> bool:
    if abs(len(a) - len(b)) > 1:
        return False

    if len(a) == len(b):
        return replace_one(a, b)

    if len(a) > len(b):
        return remove_one(b, a)

    return remove_one(a, b)


def replace_one(a: str, b: str) -> bool:
    if len(a) != len(b):
        return False

    attempts = 1
    for i in range(len(a)):
        if a[i] != b[i]:
            attempts -= 1

        if attempts < 0:
            return False

    return True


def remove_one(small: str, big: str) -> bool:
    if len(small) >= len(big):
        return False

    attempts = 1

    i = 0
    j = 0
    while i < len(small):
        if small[i] != big[j]:
            i -= 1
            attempts -= 1

        if attempts < 0:
            return False

        i += 1
        j += 1

    return i != j or attempts == 1


class Test(unittest.TestCase):
    def test_one_way(self):
        tests = [
            ['abc', 'ab', True],
            ['ab', 'abc', True],
            ['ab', 'abcc', False],
            ['pale', 'bale', True],
            ['pale', 'mape', False],
            ['', '', True],
            ['a', 'a', True],
            ['1', 'a', True],
            ['pale', 'ple', True],
            ['pales', 'pale', True],
            ['pale', 'bale', True],
            ['paleabc', 'pleabc', True],
            ['pale', 'ble', False],
            ['a', 'b', True],
            ['', 'd', True],
            ['d', 'de', True],
            ['pale', 'pale', True],
            ['pale', 'ple', True],
            ['ple', 'pale', True],
            ['pale', 'bale', True],
            ['pale', 'bake', False],
            ['pale', 'pse', False],
            ['ples', 'pales', True],
            ['pale', 'pas', False],
            ['pas', 'pale', False],
            ['pale', 'pkle', True],
            ['pkle', 'pable', False],
            ['pal', 'palks', False],
            ['palks', 'pal', False],
        ]

        for [a_str, b_str, want] in tests:
            self.assertEqual(one_way(a_str, b_str), want, f'a - {a_str}, b - {b_str}')


if __name__ == "__main__":
    unittest.main()
