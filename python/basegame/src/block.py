import pygame

from src.constants import BLOCK_SIDE

class Block:
    def __init__(self, left, top):
        self.rect = pygame.Rect(left, top, BLOCK_SIDE, BLOCK_SIDE)
        self.color = pygame.Color(100, 100, 100)

    def draw(self, window):
        pygame.draw.rect(window, self.color, self.rect)

