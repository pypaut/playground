#!/usr/bin/python3

import numpy as np


def load_input_as_matrix(n):
    matrix = np.zeros((n, n))
    with open("input.txt") as f:
        i = 0
        for line in f.readlines():
            j = 0
            for character in line:
                if character != "\n":
                    matrix[i][j] = int(character)
                    j += 1
            i += 1
    return matrix


def is_top_left_corner(i, j, n):
    return i == 0 and j == 0


def is_top_right_corner(i, j, n):
    return i == 0 and j == n


def is_bottom_left_corner(i, j, n):
    return i == n and j == 0


def is_bottom_right_corner(i, j, n):
    return i == n and j == n


def is_left_edge(i, j, n):
    return j == 0


def is_top_edge(i, j, n):
    return i == 0


def is_right_edge(i, j, n):
    return j == n


def is_bottom_edge(i, j, n):
    return i == n


def is_low_point(matrix, i, j):
    n = len(matrix) - 1
    current_val = matrix[i][j]
    if is_top_left_corner(i, j, n):
        return (
            matrix[i + 1][j] > current_val and matrix[i][j + 1] > current_val
        )
    elif is_top_right_corner(i, j, n):
        return (
            matrix[i + 1][j] > current_val and matrix[i][j - 1] > current_val
        )
    elif is_bottom_left_corner(i, j, n):
        return (
            matrix[i - 1][j] > current_val and matrix[i][j + 1] > current_val
        )
    elif is_bottom_right_corner(i, j, n):
        return (
            matrix[i - 1][j] > current_val and matrix[i][j - 1] > current_val
        )
    elif is_left_edge(i, j, n):
        return (
            matrix[i - 1][j] > current_val
            and matrix[i + 1][j]
            and matrix[i][j + 1] > current_val
        )
    elif is_top_edge(i, j, n):
        return (
            matrix[i][j - 1] > current_val
            and matrix[i][j + 1] > current_val
            and matrix[i + 1][j] > current_val
        )
    elif is_right_edge(i, j, n):
        return (
            matrix[i - 1][j] > current_val
            and matrix[i + 1][j]
            and matrix[i][j - 1] > current_val
        )
    elif is_bottom_edge(i, j, n):
        return (
            matrix[i - 1][j] > current_val
            and matrix[i][j - 1]
            and matrix[i][j + 1] > current_val
        )
    else:
        return (
            matrix[i - 1][j] > current_val
            and matrix[i + 1][j] > current_val
            and matrix[i][j - 1] > current_val
            and matrix[i][j + 1] > current_val
        )


def get_low_points_indices(matrix):
    low_points_indices = []
    n = len(matrix)
    for i in range(n):
        for j in range(n):
            if is_low_point(matrix, i, j):
                low_points_indices.append((i, j))
    return low_points_indices


def get_risk_levels(matrix, indices):
    levels = []
    for indice in indices:
        (i, j) = indice
        levels.append(matrix[i][j] + 1)
    return levels


def main():
    matrix = load_input_as_matrix(100)
    print(matrix)
    low_points_indices = get_low_points_indices(matrix)
    print(low_points_indices)
    sum_of_list_levels = sum(get_risk_levels(matrix, low_points_indices))
    print(sum_of_list_levels)


if __name__ == "__main__":
    main()
