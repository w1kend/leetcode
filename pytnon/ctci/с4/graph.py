from typing import List


class Node:
    def __init__(self, id: int):
        self.children: list[Node] = []
        self.id: int = id  # eq to a node number


def list_vertexes_to_adjacency_list(vertexes: List[List[int]], n: int, bidirectional: bool) -> List[List[int]]:
    res = [[] for _ in range(n)]

    for [source, destination] in vertexes:
        res[source].append(destination)
        if bidirectional:
            res[destination].append(source)

    return res


class Graph:
    def __init__(self, nodes: list[Node], adjacency: list[list[int]]):
        self.nodes: list[Node] = nodes
        self.__build(adjacency)

    def __build(self, adjacency: list[list[int]]):
        for i in range(len(self.nodes)):
            for node_num in adjacency[i]:
                self.nodes[i].children.append(self.nodes[node_num])
