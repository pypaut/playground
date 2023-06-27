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
    PLAYER_SIDE,
)


class Player(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()

        self.init_control_keys()
        self.load_sprites()
        self.image = pygame.surface.Surface([PLAYER_SIDE, PLAYER_SIDE])
        self.rect = self.image.get_rect(center=[W / 2, H / 2])

        # Animation
        self.ANIMATION_SPEED = ANIMATION_FPS / FPS
        self.frame_index = 0
        self.animation = "idle"
        self.frame_direction = "right"

        # Movement
        self.speed = 0.5
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

    def update_dir_with_gravity(self):
        self.direction[1] += GRAVITY_GROWTH
        if self.direction[1] > MAX_GRAVITY:
            self.direction[1] = MAX_GRAVITY

    def update_pos_with_dir(self, dt):
        self.rect.left += dt * self.speed * self.direction[0]
        self.rect.top += dt * self.speed * self.direction[1]

    def update_pos_with_collision_ground(self, blocks):
        self.is_on_ground = False
        for b in blocks:
            if self.rect.colliderect(b.rect):
                self.rect.top = b.rect.top - self.rect.height
                self.is_on_ground = True

    def update_pos_with_collision_boundaries(self):
        if self.rect.left < 0:
            self.rect.left = 0
        if self.rect.left + self.rect.width > W:
            self.rect.left = W - self.rect.width

    def update_dir_horizontal(self, keys):
        self.direction[0] = 0.0
        if keys[self.LEFT_KEY]:
            self.direction[0] -= 1.0
        if keys[self.RIGHT_KEY]:
            self.direction[0] += 1.0

    def update_dir_vertical(self, keys):
        if keys[self.JUMP_KEY] and self.is_on_ground:
            self.direction[1] = -JUMP_FORCE

    def init_control_keys(self):
        self.LEFT_KEY = pygame.K_a
        self.RIGHT_KEY = pygame.K_d
        self.JUMP_KEY = pygame.K_SPACE

    def update_animation(self):
        if self.direction[0] < 0:
            self.frame_direction = "left"
            self.animation = "run"
        elif self.direction[0] > 0:
            self.frame_direction = "right"
            self.animation = "run"
        elif self.direction[0] == 0:
            self.animation = "idle"

        self.frame_index += self.ANIMATION_SPEED
        self.frame_index %= len(
            self.sprites[self.animation][self.frame_direction]
        )
        self.image = self.sprites[self.animation][self.frame_direction][
            int(self.frame_index)
        ]

    def load_sprites(self):
        self.sprites = {}
        animations = ["idle", "run"]
        for animation in animations:
            self.sprites[animation] = {"left": [], "right": []}
            sprites_right = self.load_sprites_right(animation)
            self.sprites[animation]["right"] = sprites_right
            self.sprites[animation]["left"] = self.flip_sprites_left(
                sprites_right
            )

    def load_sprites_right(self, animation):
        right_sprites = []
        images_paths = self.get_images_paths(animation)
        for f in images_paths:
            img = pygame.image.load(f)
            img = pygame.transform.scale(img, [PLAYER_SIDE * 1.4, PLAYER_SIDE])
            right_sprites.append(img)
        return right_sprites

    def flip_sprites_left(self, right_sprites):
        sprites_left = [
            pygame.transform.flip(s.copy(), True, False) for s in right_sprites
        ]
        return sprites_left

    def get_images_paths(self, animation):
        path_animation_dir = os.path.join(ASSETS_PLAYER_PATH, animation)
        image_filenames = sorted(os.listdir(path_animation_dir))
        images_paths = [
            os.path.join(path_animation_dir, f) for f in image_filenames
        ]
        return images_paths
