import os
import pygame

from src.constants import H, W, MAX_GRAVITY, JUMP_FORCE, GRAVITY_GROWTH, PLACEHOLDER_COLOR, ASSETS_PLAYER_PATH, FPS, ANIMATION_FPS

class Player(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()

        self.init_control_keys()

        # Pygame attributes
        side = H / 7

        # Load animation images
        animation_images = {"idle": []}
        path = "assets/player"
        for animation in animation_images:
            animation_path = os.path.join(path, animation)
            filenames = sorted(os.listdir(animation_path))
            for f in filenames:
                complete_filename = os.path.join(animation_path, f)
                img = pygame.image.load(complete_filename)
                img = pygame.transform.scale(img, [side, side])
                animation_images["idle"].append(img)
        self.animation_images = animation_images

        # Load image
        self.image = self.animation_images["idle"][0]

        # Fix rect
        self.rect = self.image.get_rect(center=[W/2, H/2])

        # Animation
        self.ANIMATION_SPEED = ANIMATION_FPS / FPS
        self.current_frame = 0
        self.current_animation = "idle"

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

    def update_animation(self):
        self.current_frame += self.ANIMATION_SPEED
        self.current_frame %= len(self.animation_images[self.current_animation])
        current_frame_index = int(self.current_frame)
        self.image = self.animation_images[self.current_animation][current_frame_index]

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
