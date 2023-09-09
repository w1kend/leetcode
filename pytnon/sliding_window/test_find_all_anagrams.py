from unittest import TestCase

from l0438_find_all_anagrams import FindAllAnagrams


class TestFindAllAnagrams(TestCase):
    def test_find_anagrams(self):
        cases = [
            {'str': 'cbaebabacd', 'p': 'abc', 'want': [0, 6]},
            {'str': 'abab', 'p': 'ab', 'want': [0, 1, 2]},
            {
                'str': 'abcabc',
                'p': 'abc',
                'want': [0, 1, 2, 3],
            },
        ]

        s = FindAllAnagrams()
        for case in cases:
            self.assertEqual(case['want'], s.find_anagrams(case['str'], case['p']))
