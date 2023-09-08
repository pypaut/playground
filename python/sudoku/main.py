#!/usr/bin/python3

import numpy as np
import time


def submatrix(i, j, grid):
    """
    Get 3x3 submatrix containing the element at i, j
    """
    if 0 <= i < 3:  # Top
        if 0 <= j < 3:  # Left
            return grid[0:3, 0:3]
        elif 3 <= j < 6:  # Middle
            return grid[0:3, 3:6]
        else:  # 6 <= j < 9  # Right
            return grid[0:3, 6:9]

    elif 3 <= i < 6:  # Middle
        if 0 <= j < 3:  # Left
            return grid[3:6, 0:3]
        elif 3 <= j < 6:  # Middle
            return grid[3:6, 3:6]
        else:  # 6 <= j < 9  # Right
            return grid[3:6, 6:9]

    else:  # 6 <= i < 9  # Bottom
        if 0 <= j < 3:  # Left
            return grid[6:9, 0:3]
        elif 3 <= j < 6:  # Middle
            return grid[6:9, 3:6]
        else:  # 6 <= j < 9  # Right
            return grid[6:9, 6:9]


def solve(grid):
    """
    Actual function solving the Sudoku grid
    """
    if not 0 in grid:
        return True

    # Find first occurence of 0
    pos = np.argmax(grid == 0)
    i = pos // 9
    j = pos - 9 * i

    # Try every possible value
    for k in range(1, 10):
        row_cond = not k in grid[i]
        col_cond = not k in grid[:, j]
        submat_cond = not k in submatrix(i, j, grid)
        if row_cond and col_cond and submat_cond:
            grid[i, j] = k
            # print("###", "\n", grid)
            # time.sleep(0.1)
            if solve(grid):
                return True

    # No value was correct, reset cell
    grid[i, j] = 0
    return False


def main():
    grid = np.array(
        [
            [3, 0, 6, 5, 0, 8, 4, 0, 0],
            [5, 2, 0, 0, 0, 0, 0, 0, 0],
            [0, 8, 7, 0, 0, 0, 0, 3, 1],
            [0, 0, 3, 0, 1, 0, 0, 8, 0],
            [9, 0, 0, 8, 6, 3, 0, 0, 5],
            [0, 5, 0, 0, 9, 0, 6, 0, 0],
            [1, 3, 0, 0, 0, 0, 2, 5, 0],
            [0, 0, 0, 0, 0, 0, 0, 7, 4],
            [0, 0, 5, 2, 0, 6, 3, 0, 0],
        ]
    )

    print(solve(grid))
    print(grid)


if __name__ == "__main__":
    main()
