import pygame

from src.constants import W, MAX_GRAVITY, JUMP_FORCE, GRAVITY_GROWTH

class Player:
    def __init__(self, win_w, win_h):
        # Pygame attributes
        self.color = (255, 255, 255)
        side = win_h / 10
        self.rect = pygame.Rect(
            (win_w-side)/2, # left
            (win_h-side)/2, # top
            side,           # w
            side            # h
        )

        # Control keys
        self.left_key = pygame.K_a
        self.right_key = pygame.K_d
        self.jump_key = pygame.K_SPACE

        # Movement
        self.speed = 0.5
        self.direction = [0.0, MAX_GRAVITY]
        self.is_on_ground = False

    def events(self, keys):
        # Left and right movements
        self.direction[0] = 0.0
        if keys[self.left_key]:
            self.direction[0] = -1.0
        if keys[self.right_key]:
            self.direction[0] = 1.0
        if keys[self.jump_key] and self.is_on_ground:
            self.direction[1] = -JUMP_FORCE


    def update(self, keys, dt, blocks):
        # Gravity
        self.direction[1] += GRAVITY_GROWTH
        if self.direction[1] > MAX_GRAVITY:
            self.direction[1] = MAX_GRAVITY

        # Left and right movements
        self.rect.left += dt * self.speed * self.direction[0]
        self.rect.top += dt * self.speed * self.direction[1]

        # Vertical collisions
        self.is_on_ground = False
        for b in blocks.blocks:
            if self.rect.colliderect(b.rect):
                self.rect.top = b.rect.top - self.rect.height
                self.is_on_ground = True

        # Boundaries
        if self.rect.left < 0:
            self.rect.left = 0
        if self.rect.left + self.rect.width > W:
            self.rect.left = W - self.rect.width

    def draw(self, window):
        pygame.draw.rect(window, self.color, self.rect)
