import pygame


class Game:
    def __init__(self, width, height):
        pygame.display.init()
        pygame.font.init()
        pygame.display.set_caption("Pong")
        self.window = pygame.display.set_mode((width, height))
        self.clock = pygame.time.Clock()
        self.W, self.H = width, height
        self.is_running = False
