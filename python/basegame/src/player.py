import pygame

from src.sprites import load_player_sprites
from src.constants import (
    H,
    W,
    MAX_GRAVITY,
    JUMP_FORCE,
    GRAVITY_GROWTH,
    FPS,
    ANIMATION_FPS,
    PLAYER_WIDTH,
    PLAYER_HEIGHT,
    PLAYER_SPEED,
    TOP_OFFSET_SPRITE,
    LEFT_OFFSET_SPRITE,
    IDLE,
    RUN,
    JUMP,
    FALL,
)


class Player(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()

        # Controls
        self.LEFT_KEY = pygame.K_a
        self.RIGHT_KEY = pygame.K_d
        self.JUMP_KEY = pygame.K_SPACE

        # Collision
        self.hitrect = pygame.Rect(0, 0, PLAYER_WIDTH, PLAYER_HEIGHT)
        self.hitrect.center = [W / 2, H / 2]

        # Sprites
        animations = [IDLE, RUN, JUMP, FALL]
        self.sprites = load_player_sprites(animations)

        # Image
        self.image = self.sprites["right"][IDLE][0]
        self.rect = self.image.get_rect(center=self.hitrect.center)
        self.rect.top -= TOP_OFFSET_SPRITE
        self.rect.left += LEFT_OFFSET_SPRITE

        # Animation
        self.ANIMATION_SPEED = ANIMATION_FPS / FPS
        self.frame_index = 0
        self.animation = IDLE
        self.frame_direction = "right"

        # Movement
        self.speed = PLAYER_SPEED
        self.direction = [0.0, MAX_GRAVITY]
        self.is_on_ground = False

    def events(self, keys):
        """
        Update direction according to input
        """
        self.update_dir_horizontal(keys)
        self.update_dir_vertical(keys)

    def update(self, keys, dt, blocks):
        """
        Update position according to direction and collision
        """
        self.update_dir_with_gravity()
        self.update_pos_with_dir_y(dt)
        self.update_pos_with_collision_y(blocks)
        old_x = self.hitrect.x
        self.update_pos_with_dir_x(dt)
        self.update_pos_with_collision_x(blocks, old_x)
        self.update_pos_with_collision_boundaries()
        self.update_animation()
        self.update_sprite_rect()

    def update_sprite_rect(self):
        self.rect.center = self.hitrect.center  # center[0] != left !!!
        self.rect.top -= TOP_OFFSET_SPRITE
        if self.frame_direction == "left":
            self.rect.left -= LEFT_OFFSET_SPRITE
        else:
            self.rect.left += LEFT_OFFSET_SPRITE

    def update_dir_with_gravity(self):
        self.direction[1] += GRAVITY_GROWTH
        if self.direction[1] > MAX_GRAVITY:
            self.direction[1] = MAX_GRAVITY

    def update_pos_with_dir_x(self, dt):
        self.hitrect.left += dt * self.speed * self.direction[0]

    def update_pos_with_dir_y(self, dt):
        self.hitrect.top += dt * self.speed * self.direction[1]

    def update_pos_with_collision_y(self, blocks):
        self.is_on_ground = False
        for b in blocks:
            if self.hitrect.colliderect(b.rect):
                if self.direction[1] > 0:
                    self.hitrect.top = b.rect.top - self.hitrect.height
                    self.is_on_ground = True
                elif self.direction[1] < 0:
                    self.hitrect.top = b.rect.top + b.rect.height
                    self.direction[1] = 0

    def update_pos_with_collision_x(self, blocks, old_x):
        for b in blocks:
            if self.hitrect.colliderect(b.rect):
                self.hitrect.x = old_x

    def update_pos_with_collision_boundaries(self):
        if self.hitrect.left < 0:
            self.hitrect.left = 0
        if self.hitrect.left + self.hitrect.width > W:
            self.hitrect.left = W - self.hitrect.width

    def update_dir_horizontal(self, keys):
        self.direction[0] = 0.0
        if keys[self.LEFT_KEY]:
            self.direction[0] -= 1.0
        if keys[self.RIGHT_KEY]:
            self.direction[0] += 1.0

    def update_dir_vertical(self, keys):
        if keys[self.JUMP_KEY] and self.is_on_ground:
            self.direction[1] = -JUMP_FORCE

    def update_animation(self):
        if self.direction[0] < 0:
            self.frame_direction = "left"
            self.animation = RUN
        elif self.direction[0] > 0:
            self.frame_direction = "right"
            self.animation = RUN
        elif self.direction[0] == 0:
            self.animation = IDLE
        if self.direction[1] < 0 and not self.is_on_ground:
            self.animation = JUMP
        elif self.direction[1] > 0 and not self.is_on_ground:
            self.animation = FALL

        self.frame_index += self.ANIMATION_SPEED
        self.frame_index %= len(
            self.sprites[self.frame_direction][self.animation]
        )
        self.image = self.sprites[self.frame_direction][self.animation][
            int(self.frame_index)
        ]
