import unittest
from typing import List


# Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0
def zero_matrix(m: List[List[int]]) -> List[List[int]]:
    # O(NxM)
    zero_rows = dict()
    zero_cols = dict()

    for row in range(len(m)):
        for col in range(len(m[row])):
            if m[row][col] == 0:
                zero_cols[col] = 1
                zero_rows[row] = 1

    for row in zero_rows:
        zero_row(m, row)

    for row in range(len(m)):
        # maybe the row is already zero
        if row not in zero_rows:
            for col in zero_cols:
                m[row][col] = 0

    return m


def zero_row(matrix: List[List[int]], row: int):
    for col in range(len(matrix[row])):
        matrix[row][col] = 0


class Test(unittest.TestCase):
    def test_zero_matrix(self):
        tests = [
            (
                [
                    [1, 2, 3],
                    [1, 0, 0],
                    [2, 3, 4],
                ],
                [
                    [1, 0, 0],
                    [0, 0, 0],
                    [2, 0, 0],
                ]
            ),
            (
                [
                    [1, 1, 2, 4, 1, 2],
                    [2, 3, 0, 4, 1, 5],
                    [1, 3, 4, 2, 0, 3],
                    [1, 2, 3, 4, 4, 3],
                ],
                [
                    [1, 1, 0, 4, 0, 2],
                    [0, 0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0, 0],
                    [1, 2, 0, 4, 0, 3],
                ]
            ),
            (
                [
                    [1, 1],
                    [2, 3],
                ],
                [
                    [1, 1],
                    [2, 3],
                ],
            ),
            (
                [
                    [0, 1],
                    [2, 0],
                ],
                [
                    [0, 0],
                    [0, 0],
                ],
            ),
            (
                [
                    [1],
                ],
                [
                    [1],
                ],
            ),
            (
                [
                    [1, 2, 3, 4, 0],
                    [6, 0, 8, 9, 10],
                    [11, 12, 13, 14, 15],
                    [16, 0, 18, 19, 20],
                    [21, 22, 23, 24, 25]
                ],
                [
                    [0, 0, 0, 0, 0],
                    [0, 0, 0, 0, 0],
                    [11, 0, 13, 14, 0],
                    [0, 0, 0, 0, 0],
                    [21, 0, 23, 24, 0]
                ]
            )
        ]

        for [matrix, want] in tests:
            self.assertEqual(want, zero_matrix(matrix))


if __name__ == '__main__':
    unittest.main()
