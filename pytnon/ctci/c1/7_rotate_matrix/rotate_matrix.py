import unittest
from typing import List


def rotate_new(m: List[List[int]]) -> List[List[int]]:
    size = len(m)
    # create an empty matrix
    res = [[0] * size for _ in range(size)]

    for i in range(size):
        for j in range(size):
            res[j][size - i - 1] = m[i][j]

    return res


def rotate_in_place(m: List[List[int]]) -> List[List[int]]:
    size = len(m)

    for layer in range(size // 2):
        for i in range(layer, size - layer - 1):
            top = m[layer][i]

            # move left side to top
            m[layer][i] = m[size - i - 1][layer]

            # move bottom to the left
            m[size - i - 1][layer] = m[-layer - 1][size - i - 1]

            # move right to the bottom
            m[-layer - 1][size - i - 1] = m[i][-layer - 1]

            # move the last
            m[i][-layer - 1] = top

    return m


class Test(unittest.TestCase):
    def test_rotate_matrix(self):
        data = [
            {
                'matrix': [
                    [1, 2, 3],
                    [4, 5, 6],
                    [7, 8, 9]
                ],
                'want': [
                    [7, 4, 1],
                    [8, 5, 2],
                    [9, 6, 3]
                ]
            },
            {
                'matrix': [
                    [1, 2, 3, 4],
                    [5, 6, 7, 8],
                    [9, 10, 11, 12],
                    [13, 14, 15, 16]
                ],
                'want': [
                    [13, 9, 5, 1],
                    [14, 10, 6, 2],
                    [15, 11, 7, 3],
                    [16, 12, 8, 4]
                ],
            },
            {
                'matrix': [
                    [1, 2, 3, 4, 5],
                    [6, 7, 8, 9, 10],
                    [11, 12, 13, 14, 15],
                    [16, 17, 18, 19, 20],
                    [21, 22, 23, 24, 25]
                ],
                'want': [
                    [21, 16, 11, 6, 1],
                    [22, 17, 12, 7, 2],
                    [23, 18, 13, 8, 3],
                    [24, 19, 14, 9, 4],
                    [25, 20, 15, 10, 5]
                ],
            }
        ]

        for case in data:
            self.assertEqual(case['want'], rotate_new(case['matrix']))
            self.assertEqual(case['want'], rotate_in_place(case['matrix']))
