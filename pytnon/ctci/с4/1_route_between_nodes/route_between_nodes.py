import unittest

from pytnon.ctci.с4.graph import Graph, Node, list_vertexes_to_adjacency_list


def route_between_nodes(graph: Graph, source: int, destination: int) -> bool:
    # BFS
    seen = set()
    queue = [graph.nodes[source]]

    while queue:
        node = queue.pop()
        if node.id == destination:
            return True
        seen.add(node.id)
        for child in node.children:
            if child.id not in seen:
                queue.append(child)

    return False


class Test(unittest.TestCase):
    def test_route_exists_between_nodes(self):
        tests = [
            (
                Graph(
                    [Node(i) for i in range(3)],
                    [[1, 2], [0, 2], [0, 1]]
                ),
                0,
                2,
                True,
            ),
            (
                Graph(
                    [Node(i) for i in range(6)],
                    [[1, 2], [0], [0], [4, 5], [3, 5], [3, 4]]
                ),
                0,
                5,
                False,
            ),
            (
                Graph(
                    [Node(i) for i in range(6)],
                    list_vertexes_to_adjacency_list([[0, 1], [0, 2], [3, 5], [5, 4], [4, 3]], 6, True)
                ),
                0,
                5,
                False,
            ),
            (
                Graph(
                    [Node(i) for i in range(6)],
                    [[1, 2], [0], [0], [4, 5], [3, 5], [3, 4]]
                ),
                0,
                1,
                True,
            )
        ]

        for [graph, source, destination, expected] in tests:
            self.assertEqual(expected, route_between_nodes(graph, source, destination))
