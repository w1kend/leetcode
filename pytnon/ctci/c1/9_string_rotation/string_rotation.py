# Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings,
# s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring
# (e.g.,"waterbottle" is a rotation of"erbottlewat").
import unittest


def is_rotated_string(s1: str, s2: str) -> bool:
    if len(s1) == len(s2) != 0:
        return s1 in s2 + s2
    return False


class Test(unittest.TestCase):
    def test_is_rotated_string(self):
        tests = [
            ('waterbottle', 'erbottlewat', True),
            ('foo', 'bar', False),
            ('foo', 'foofoo', False),
            ('foo', 'oof', True)
        ]

        for [s1, s2, want] in tests:
            self.assertEqual(want, is_rotated_string(s1, s2))


if __name__ == '__main__':
    unittest.main()
