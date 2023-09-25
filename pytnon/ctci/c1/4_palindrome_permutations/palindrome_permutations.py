import unittest
from collections import defaultdict


def palindrome_perms(word: str) -> bool:
    # trim spaces
    word = word.replace(' ', '')

    # use dict to count chars in the word
    dic = defaultdict(int)

    for c in word:
        dic[c.lower()] += 1

    # for even chars number all the characters in the word should be presented an even number of times
    is_even = len(word) % 2 == 0

    max_odd_chars_num = 0 if is_even else 1

    odd_nums = 0
    for char, times in dic.items():
        if times % 2 != 0:
            odd_nums += 1

        if odd_nums > max_odd_chars_num:
            return False

    return True


class Test(unittest.TestCase):
    def test_palindrome_permutations(self):
        tests = [
            ['Tact Coa', True],
            ['ab', False],
            ['a   b a', True],
            ['jhsabckuj ahjsbckj', True],
            ['Able was I ere I saw Elba', True],
            ['So patient a nurse to nurse a patient so', False],
            ['Random Word', False],
            ['no x in nixon', True]
        ]

        for word, want in tests:
            self.assertEqual(want, palindrome_perms(word))
