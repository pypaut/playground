import pygame


class Ball:
    def __init__(self, rect, color):
        self.rect = rect
        self.speed = 1
        self.color = color

    def update(self, dt):
        pass

    def draw(self, window):
        pygame.draw.rect(window, self.color, self.rect)

