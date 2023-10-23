import unittest
from typing import List


def build_order(projects: List[str], dependencies: List[tuple[str, str]]) -> List[str]:
    project_deps: dict[str, set[str]] = {proj: set() for proj in projects}

    for dependency, project in dependencies:
        project_deps[project].add(dependency)

    build_sequence: List[str] = []
    to_build = {proj: "" for proj in projects}  # dict is ordered, so we can have results with the same order
    while to_build:
        built = False
        for project in list(to_build):
            if not project_deps[project].intersection(to_build):
                del to_build[project]
                build_sequence.append(project)
                built = True

        # if we didn't build anything than dependencies are cycled
        if not built:
            return []

    return build_sequence


class Test(unittest.TestCase):
    def test_build_order(self):
        tests = [
            (
                ["a", "b", "c", "d", "e", "f"],
                [
                    ('a', 'd'),
                    ('f', 'b'),
                    ('b', 'd'),
                    ('f', 'a'),
                    ('d', 'c'),
                ],
                ['e', 'f', 'a', 'b', 'd', 'c']
            ),
            (
                ["a", "b", "c", "d", "e", "f", "g"],
                [
                    ("d", "g"),
                    ("a", "e"),
                    ("b", "e"),
                    ("c", "a"),
                    ("f", "a"),
                    ("b", "a"),
                    ("f", "c"),
                    ("f", "b"),
                ],
                ['d', 'f', 'g', 'b', 'c', 'a', 'e'],
            )
        ]

        for projects, dependencies, expected in tests:
            self.assertEqual(expected, build_order(projects, dependencies))
