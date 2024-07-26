import configparser
import json
import pygame
import socket

from pygame.locals import QUIT, KEYDOWN


class Client:
    def __init__(self):
        # Read config
        config = configparser.ConfigParser()
        with open("config.ini", "r") as configfile:
            config.read_file(configfile)

        self.H = config.getint("CLIENT", "H")
        self.W = config.getint("CLIENT", "W")
        self.FPS = config.getint("CLIENT", "FPS")

        # Init PyGame
        pygame.display.init()
        pygame.font.init()
        pygame.key.set_repeat(1, 1)
        pygame.display.set_caption("Multiplayer")

        self.window = pygame.display.set_mode((self.W, self.H))
        self.clock = pygame.time.Clock()

        # Network
        self.socket = socket.socket()
        self.socket.connect(("", 12350))

    def run(self):
        while True:
            self.clock.tick(self.FPS)

            # Quit event
            events = pygame.event.get()
            if QUIT in [e.type for e in events]:
                break

            # Keyboard events
            if KEYDOWN in [e.type for e in events]:
                keys = pygame.key.get_pressed()
                if keys[pygame.K_ESCAPE]:
                    break

                player_dir = [0, 0]

                if keys[pygame.K_w]:
                    player_dir[1] -= 1
                if keys[pygame.K_a]:
                    player_dir[0] -= 1
                if keys[pygame.K_s]:
                    player_dir[1] += 1
                if keys[pygame.K_d]:
                    player_dir[0] += 1

                if player_dir != [0, 0]:
                    player_dir = json.dumps({"x": player_dir[0], "y": player_dir[1]})
                    self.socket.send(f"{json.dumps(player_dir)}".encode("utf-8"))
