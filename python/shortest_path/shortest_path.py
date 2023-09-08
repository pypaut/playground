#!/usr/bin/python3

"""
The goal of this script is to generate the most efficient
route through each node, with some nodes to be reached before
the others
"""

import itertools
import numpy as np


def FloydWarshall(vertices, edges, weights):
    nb_vertices = len(vertices)
    nb_edges = len(edges)
    nb_weights = len(weights)
    if nb_edges != nb_weights:
        print("ERROR : nb_edges != nb_weights")
        exit(1)

    distances = np.full((nb_vertices, nb_vertices), 999999)
    for i in range(nb_vertices):
        distances[i][i] = 0

    for i in range(nb_edges):
        e = edges[i]
        w = weights[i]
        distances[e[0]][e[1]] = w

    for k in range(nb_vertices):
        for i in range(nb_vertices):
            for j in range(nb_vertices):
                if distances[i][j] > distances[i][k] + distances[k][j]:
                    distances[i][j] = distances[i][k] + distances[k][j]

    return distances


def ComputePermutations(vertices):
    return list(itertools.permutations(vertices))


def PermutationVerifiesConstraints(permutation, constraints):
    for i in range(len(permutation)):
        if constraints[permutation[i]] != None:
            for j in range(len(permutation[:i])):
                if permutation[j] == constraints[permutation[i]]:
                    return True
    return False


def BruteForceShortestPathCost(vertices, constraints, edges, weights):
    answer_cost = 999999
    answer_permutation = None
    distances = FloydWarshall(vertices, edges, weights)
    permutations = ComputePermutations(vertices)

    for permutation in permutations:
        if not PermutationVerifiesConstraints(permutation, constraints):
            continue
        cost = 0
        previous = permutation[0]
        for node in permutation:
            cost += distances[previous][node]
            previous = node
        if cost < answer_cost:
            answer_cost = cost
            answer_permutation = permutation

    return answer_cost, answer_permutation


def main():
    vertices = [0, 1, 2, 3]
    constraints = [None, 3, None, None]
    edges = [(0, 1), (0, 2), (2, 3), (3, 2), (1, 2), (2, 1)]
    weights = [1, 1, 1, 2, 5, 3]

    names = [
        "ZAAP VILLAGE",
        "SRAM",
        "ENUTROF",
        "ÉCAFLIP",
        "ZOBAL",
        "SACRIEUR",
        "PANDAWA",
        "XÉLOR",
        "ROUBLARD",
        "IOP",
        "CRÂ",
        "SADIDA",
        "ENIRIPSA",
        "OSAMODAS",
        "STEAMER",
        "FECA",
        "VILE GREDINE",
    ]

    vertices = [i for i in range(len(names))]

    edges = [
        (names.index("ZAAP ALMANAX"), names.index("ALMANAX")),
        ()
    ]

    weights = [
        1,
        1,
    ]

    cost, order = BruteForceShortestPathCost(vertices, constraints, edges, weights)
    print(cost)


if __name__ == "__main__":
    main()
