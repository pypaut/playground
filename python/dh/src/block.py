import pygame


class Block:
    def __init__(self, x, y, w, h):
        self.rect = pygame.Rect(x, y, w, h)
        self.color = (200, 200, 200)

    def draw(self, window):
        """
        Draw block on window
        """
        pygame.draw.rect(window, self.color, self.rect)
