import pygame

from src.constants import BLOCK_SIDE


class Block(pygame.sprite.Sprite):
    def __init__(self, left, top):
        super().__init__()

        # Load image
        image = pygame.image.load("assets/block_tile.png")
        self.image = pygame.transform.scale(image, [BLOCK_SIDE, BLOCK_SIDE])

        # Fix rect
        self.rect = self.image.get_rect()
        self.rect.left = left
        self.rect.top = top
