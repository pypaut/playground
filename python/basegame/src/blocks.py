from src.block import Block
from src.constants import BLOCK_SIDE, W, H

class Blocks:
    def __init__(self):
        self.blocks = []

        # Ground
        nb_blocks_in_width = int(W // BLOCK_SIDE) + 1
        ground_height = H - BLOCK_SIDE
        ground_blocks = [
            Block(
                i * BLOCK_SIDE,
                ground_height,
            ) for i in range(nb_blocks_in_width)
        ]

        self.blocks += ground_blocks

    def draw(self, window):
        for b in self.blocks:
            b.draw(window)
