import unittest
from typing import List, Optional


class Node:
    def __init__(self, name: str):
        self.children: list[Node] = []
        self.name = name


class Graph:
    def __init__(self, nodes: list[Node], adjacency: List[tuple[str, str]]):
        self.nodes: dict[str, Node] = {node.name: node for node in nodes}
        self.__build(adjacency)

    def __build(self, adjacency: List[tuple[str, str]]):
        for start, finish in adjacency:
            self.nodes[start].children.append(self.nodes[finish])


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


def build_order_dfs(projects: List[str], dependencies: List[tuple[str, str]]) -> List[str]:
    # build a graph
    nodes: List[Node] = [Node(project) for project in projects]
    adjacency_list: List[tuple[str, str]] = [(dependency, project) for dependency, project in dependencies]
    graph: Graph = Graph(nodes, adjacency_list)

    created_projects: set[str] = set()
    creation_sequence: List[str] = []

    for project, node in graph.nodes.items():
        if project not in created_projects:
            dfs(node, created_projects, creation_sequence)

    creation_sequence.reverse()
    return creation_sequence


def dfs(node: Optional[Node], created: set[str], sequence: List[str]):
    # create project if it doesn't have any dependencies
    if len(node.children) == 0:
        created.add(node.name)
        sequence.append(node.name)
        return

    for child in node.children:
        # create all dependencies first
        if child.name not in created:
            dfs(child, created, sequence)
    # create current project
    created.add(node.name)
    sequence.append(node.name)


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
                ['e', 'f', 'a', 'b', 'd', 'c'],
                ['f', 'e', 'b', 'a', 'd', 'c']
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
                ['f', 'd', 'g', 'c', 'b', 'a', 'e'],
            )
        ]

        for projects, dependencies, expected, graph_expected in tests:
            self.assertEqual(expected, build_order(projects, dependencies))
            self.assertEqual(graph_expected, build_order_dfs(projects, dependencies))
