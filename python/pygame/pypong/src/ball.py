import pygame


class Ball:
    def __init__(self, rect, color, boundary):
        self.rect = rect
        self.speed = 1
        self.color = color
        self.dir_x = 0.0
        self.dir_y = 0.0
        self.boundary = boundary

    def update(self, dt):
        self.rect.top += self.dir_x * dt
        self.rect.left += self.dir_y * dt
        return self.boundary.contains(self.rect)

    def draw(self, window):
        pygame.draw.rect(window, self.color, self.rect)
