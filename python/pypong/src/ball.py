import pygame


class Ball:
    def __init__(self, rect, color, boundary):
        self.rect = rect
        self.speed = 0.4
        self.color = color
        self.dir_x = 0.0
        self.dir_y = 0.0
        self.boundary = boundary

    def update(self, p1, p2, dt):
        self.handle_player_collision(p1, p2, dt)

        if not self.handle_wall_collision():
            return False

        self.update_position(dt)
        return True

    def update_position(self, dt):
        self.rect.left += self.dir_x * dt
        self.rect.top += self.dir_y * dt

    def handle_player_collision(self, p1, p2, dt):
        p1_collides = self.rect.colliderect(p1.rect)
        p2_collides = self.rect.colliderect(p2.rect)

        if p1_collides or p2_collides:
            self.dir_x *= -1
            self.dir_y *= -1

            ball_middle_y = self.rect.top + self.rect.height // 2

            if p1_collides:
                player_middle_y = p1.rect.top + p1.rect.height // 2
            else:
                player_middle_y = p2.rect.top + p2.rect.height // 2

            player_collision_y = ball_middle_y - player_middle_y
            self.dir_y += 0.005 * player_collision_y

            self.dir_x, self.dir_y = self.get_normalized_dir()

    def handle_wall_collision(self):
        if self.rect.top < 0 or self.rect.top + self.rect.height > self.boundary.height:
            self.dir_y *= -1

        elif self.rect.left < 0 or self.rect.left + self.rect.width > self.boundary.width:
            return False

        return True


    def get_normalized_dir(self):
        norm = (self.dir_x**2 + self.dir_y**2) ** (1 / 2)
        return self.dir_x / norm * self.speed, self.dir_y / norm * self.speed

    def draw(self, window):
        pygame.draw.rect(window, self.color, self.rect)
