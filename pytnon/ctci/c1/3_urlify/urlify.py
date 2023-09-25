import unittest
from io import StringIO


def urlify(url: str, length: int) -> str:
    str_builder = StringIO()

    for i in range(length):
        if url[i] == ' ':
            str_builder.write('%20')
        else:
            str_builder.write(url[i])

    return str_builder.getvalue()


class Test(unittest.TestCase):
    def test_urlify(self):
        data = [
            ['Mr John Smith ', 13, 'Mr%20John%20Smith'],
            [' ', 1, '%20'],
            ['abc    12', 9, 'abc%20%20%20%2012'],
            ['much ado about nothing      ', 22, 'much%20ado%20about%20nothing']
        ]

        for url, length, expected in data:
            self.assertEqual(urlify(url, length), expected)


if __name__ == '__main__':
    unittest.main()
