import pygame

from src.constants import W, H
from src.sprite_functions import load_sprites, left_sprites


class Player:
    def __init__(self):
        """
        Load sprites and setup initial state
        """
        # Some options
        self.NB_FRAMES_ANIM = 5
        self.SPEED = 0.4
        self.GRAVITY = 0.8
        self.JUMP_FORCE = 0.8

        # Load sprites
        self.sprites = {
            "right": {
                "idle": load_sprites("media/character/idle"),
                "run": load_sprites("media/character/Run"),
                "attack": load_sprites("media/character/Attack"),
                "jump": load_sprites("media/character/Jump"),
                "fall": load_sprites("media/character/Fall"),
            },
        }
        self.sprites["left"] = {
            a: left_sprites(s) for (a, s) in self.sprites["right"].items()
        }

        # Initial state
        self.direction = "right"
        self.animation = "idle"
        self.is_attacking = False
        self.is_jumping = False
        self.is_falling = False
        h, w = self.sprites["right"]["idle"][0].get_size()
        self.rect = pygame.Rect(H // 2, W // 2, h, w)
        self.momentum = self.GRAVITY

        # Variables
        self.frame_counter = 0
        self.index_sprite = 0
        self.last_anim = self.animation

    def update_state(self, keys, dt):
        """
        Handle input and update state
        """
        # Movements
        if keys[pygame.K_a] and not keys[pygame.K_d]:
            self.direction = "left"
            self.animation = "run"
        elif keys[pygame.K_d] and not keys[pygame.K_a]:
            self.direction = "right"
            self.animation = "run"
        else:
            self.animation = "idle"

        # Jump
        if keys[pygame.K_SPACE] and not self.is_jumping:
            self.is_jumping = True
            self.momentum = -self.JUMP_FORCE
            self.animation = "jump"

        if self.is_jumping:
            self.animation = "jump"

        # Fall
        if self.is_falling:
            self.animation = "fall"

        # Attack
        if keys[pygame.K_q] and not self.is_attacking:
            self.is_attacking = True
            self.animation = "attack"

        if self.is_attacking:
            self.animation = "attack"

    def update_sprite(self):
        """
        Update current sprite index
        """
        # Reset if new animation
        if self.last_anim != self.animation:
            self.index_sprite = 0
            self.frame_counter = 0
            self.last_anim = self.animation

        # Update indices
        if self.frame_counter >= self.NB_FRAMES_ANIM - 1:
            nb_sprites = len(self.sprites[self.direction][self.animation])
            self.index_sprite += 1
            # End of attack
            if self.is_attacking and self.index_sprite == nb_sprites - 1:
                self.animation = "idle"
                self.is_attacking = False
            self.frame_counter = 0
            self.index_sprite %= nb_sprites - 1
        else:
            self.frame_counter += 1

    def apply_y_movement(self, dt):
        """
        Apply gravity and jump
        """
        if not self.is_attacking:
            self.rect = self.rect.move(0, self.momentum * dt)

        self.momentum += 0.05
        if self.momentum > self.GRAVITY:
            self.momentum = self.GRAVITY

    def apply_x_movement(self, keys, dt):
        """
        Apply left/right movement
        """
        if (
            keys[pygame.K_a] or keys[pygame.K_d]
        ) and not self.is_attacking:
            if self.direction == "left" and self.rect.x > 10:
                self.rect = self.rect.move(-self.SPEED * dt, 0)
            elif (
                self.direction == "right"
                and self.rect.x < W - self.get_sprite().get_size()[0] - 10
            ):
                self.rect = self.rect.move(self.SPEED * dt, 0)

    def check_x_collision(self, blocks):
        """
        Check collision on X axis
        """
        for b in blocks:
            if self.rect.colliderect(b.rect):
                offset = 0
                if b.rect.x > self.rect.x:  # Block is on left
                    offset = self.rect.x + self.rect.w - b.rect.x
                if b.rect.x < self.rect.x:  # Block is on right
                    offset = self.rect.x - (b.rect.x + b.rect.w)
                self.rect = self.rect.move(-offset, 0)

    def check_y_collision(self, blocks):
        """
        Check collision on Y axis
        """
        has_collided = False
        for b in blocks:
            if self.rect.colliderect(b.rect):
                has_collided = True
                offset = 0
                if b.rect.y > self.rect.y:  # Block is under
                    offset = self.rect.y + self.rect.h - b.rect.y
                if b.rect.y < self.rect.y:  # Block is above
                    offset = self.rect.y - (b.rect.y + b.rect.h)
                self.rect = self.rect.move(0, -offset)
        if has_collided:
            self.is_falling = False
            self.is_jumping = False
        else:
            self.is_falling = True

    def update_pos(self, keys, blocks, dt):
        """
        Update player position according to current movement
        """
        self.apply_y_movement(dt)
        self.check_y_collision(blocks)
        self.apply_x_movement(keys, dt)
        self.check_x_collision(blocks)

    def update(self, keys, blocks, dt):
        """
        Global update according to input
        """
        self.update_state(keys, dt)
        self.update_sprite()
        self.update_pos(keys, blocks, dt)

    def get_sprite(self):
        """
        Getter for current sprite surface, according to direction, animation
        and index
        """
        return self.sprites[self.direction][self.animation][self.index_sprite]
