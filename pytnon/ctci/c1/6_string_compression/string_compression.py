import unittest
from io import StringIO


def string_compression(s: str) -> str:
    if len(s) < 3:
        return s

    current = s[0]
    cnt = 0

    # aaabbbacc
    # a3b3a1c2
    result = StringIO()
    for c in s:
        if c == current:
            cnt += 1
            continue

        result.write(current + str(cnt))
        current = c
        cnt = 1

    result.write(current + str(cnt))

    result = result.getvalue()

    return min(s, result, key=len)


class Test(unittest.TestCase):
    def test_string_compression(self):
        tests = [
            ('aabcccccaaa', 'a2b1c5a3'),
            ('abcdef', 'abcdef'),
            ('aabbcc', 'aabbcc'),
            ('aaa', 'a3'),
            ('', ''),
            ('ab', 'ab'),
            ('aab', 'aab'),
        ]

        for [s, want] in tests:
            self.assertEqual(string_compression(s), want, f'str - {s}, want - {want}')


if __name__ == "__main__":
    unittest.main()
