import os
import pygame

from src.constants import (
    H,
    W,
    MAX_GRAVITY,
    JUMP_FORCE,
    GRAVITY_GROWTH,
    ASSETS_PLAYER_PATH,
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
        self.init_control_keys()

        # Collision
        self.hitrect = pygame.Rect(0, 0, PLAYER_WIDTH, PLAYER_HEIGHT)
        self.hitrect.center = [W / 2, H / 2]

        # Sprites
        animations = [IDLE, RUN, JUMP, FALL]
        self.init_load_sprites(animations)

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
        self.update_pos_with_dir(dt)
        self.update_pos_with_collision_ground(blocks)
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

    def update_pos_with_dir(self, dt):
        x = dt * self.speed * self.direction[0]
        y = dt * self.speed * self.direction[1]
        self.move(x, y)

    def update_pos_with_collision_ground(self, blocks):
        self.is_on_ground = False
        for b in blocks:
            if self.hitrect.colliderect(b.rect):
                # Collides, thus compute offset with floor and move up
                target_top = b.rect.top - self.hitrect.height
                y_translation = target_top - self.hitrect.top
                self.move(0, y_translation)
                self.is_on_ground = True

    def update_pos_with_collision_boundaries(self):
        if self.hitrect.left < 0:
            self.move(-self.hitrect.left, 0)  # Reset to 0
        if self.hitrect.left + self.hitrect.width > W:
            x_translation = -(self.hitrect.left + self.hitrect.width - W)
            self.move(x_translation, 0)

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

    def init_control_keys(self):
        self.LEFT_KEY = pygame.K_a
        self.RIGHT_KEY = pygame.K_d
        self.JUMP_KEY = pygame.K_SPACE

    def init_load_sprites(self, animations):
        self.sprites = {
            "right": {
                a: self.load_sprites(os.path.join(ASSETS_PLAYER_PATH, a))
                for a in animations
            },
        }
        self.sprites["left"] = {
            a: self.left_sprites(s) for (a, s) in self.sprites["right"].items()
        }

    def load_sprites(self, path):
        filenames = sorted(os.listdir(path))
        sprites = [pygame.image.load(os.path.join(path, f)) for f in filenames]
        sprites = list(map(self.scale_sprite, sprites))
        return sprites

    def scale_sprite(self, sprite):
        return pygame.transform.scale(
            sprite, [PLAYER_WIDTH * 3.5, PLAYER_HEIGHT * 1.3]
        )

    def left_sprites(self, sprites):
        return [pygame.transform.flip(i.copy(), True, False) for i in sprites]

    def move(self, x, y):
        self.rect.left += x
        self.rect.top += y
        self.hitrect.left += x
        self.hitrect.top += y
