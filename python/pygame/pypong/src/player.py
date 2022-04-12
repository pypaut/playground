import pygame

from enum import Enum


class ControlType(Enum):
    PLAYER1 = 0
    PLAYER2 = 1


class Player:
    def __init__(self, rect, controls):
        self.rect = rect
        self.speed = 1
        self.set_control_keys(controls)

    def update(self, keys, dt, w, h):
        if keys[self.K_UP]:
            self.rect.top -= self.speed * dt

        if keys[self.K_DOWN]:
            self.rect.top += self.speed * dt

        self.rect = self.rect.clamp(pygame.Rect(0, 0, w, h))

    def set_control_keys(self, controls):
        if controls == ControlType.PLAYER1:
            self.K_UP = pygame.K_w
            self.K_DOWN = pygame.K_s
        elif controls == ControlType.PLAYER2:
            self.K_UP = pygame.K_UP
            self.K_DOWN = pygame.K_DOWN
