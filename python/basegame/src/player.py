import pygame

from src.constants import H, W, MAX_GRAVITY, JUMP_FORCE, GRAVITY_GROWTH, PLACEHOLDER_COLOR

class Player(pygame.sprite.Sprite):
    def __init__(self):
        super().__init__()

        # Pygame attributes
        side = H / 7

        # Load image
        image = pygame.image.load("assets/character/idle/Warrior_Idle_1.png")
        self.image = pygame.transform.scale(image, [side, side])

        self.rect = self.image.get_rect()
        self.rect.center = [W/2, H/2]

        self.init_control_keys()

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
