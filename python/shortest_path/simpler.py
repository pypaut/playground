#!/usr/bin/python3

"""
The goal of this script is to generate the most efficient
route through each node
"""

import itertools
import math as m
import numpy as np


def distance(p1, p2):
    return math.abs(p2[0] - p1[0]) + math.abs(p2[1] - p1[1])


def compute_cost(permutation):
    cost = 0
    for i in range(len(permutation)) - 1:
        cost += distance(permutation[i][1], permutation[i + 1][1])
    return cost


def main():
    nodes = [
        ("Zaap", (-2, 0)),
        ("Captain", (0, 0)),
        ("Vile gredine", (-1, 0)),
        ("Sram", (-4, 0)),
        ("Enutrof", (-1, 4)),
        ("Ecaflip", (1, -5)),
        ("Zobal", (1, -8)),
        ("Sacrieur", (-2, -8)),
        ("Pandawa", (4, -2)),
        ("Crâ", (0, 3)),
        ("Iop", (1, 3)),
        ("Roublard", (3, 3)),
        ("Xélor", (3, 1)),
        ("Sadida", (-1, 9)),
        ("Eniripsa", (7, 1)),
        ("Steamer", (9, 1)),
        ("Osamodas", (8, 2)),
        ("Feca", (12, 5)),
    ]

    print("Compute permutations...")
    permutations = list(itertools.permutations(nodes))
    costs = []

    print("Compute costs...")
    for permutation in permutations:
        costs.append(compute_cost(permutation))

    min_cost = min(costs)
    min_index = costs.index(min_cost)

    print(permutations(min_index))


if __name__ == "__main__":
    main()
