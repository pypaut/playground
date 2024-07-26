import configparser
import json
import math
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
        print("Connected to server")

        # Game engine
        self.player_pos = {"x": 0, "y": 0}

    def run(self):
        while True:
            self.clock.tick(self.FPS)

            # Quit event
            events = pygame.event.get()
            if QUIT in [e.type for e in events]:
                break

            # Keyboard events
            keys = pygame.key.get_pressed()
            if keys[pygame.K_ESCAPE]:
                break

            player_dir = {"x": 0., "y": 0.}

            if keys[pygame.K_w]:
                player_dir["y"] -= 1
            if keys[pygame.K_a]:
                player_dir["x"] -= 1
            if keys[pygame.K_s]:
                player_dir["y"] += 1
            if keys[pygame.K_d]:
                player_dir["x"] += 1

            norm = math.sqrt(sum([x**2 for x in player_dir.values()]))
            if norm:
                player_dir["x"] /= norm
                player_dir["y"] /= norm

            # Send direction
            # if player_dir["x"] or player_dir["y"]:
            self.socket.send(f"{json.dumps(player_dir)}".encode("utf-8"))

            # Receive position
            server_msg = self.socket.recv(1024)
            if server_msg != b"":
                self.player_pos = json.loads(server_msg)

            # Draw
            self.window.fill((0, 0, 0))
            pygame.draw.rect(
                self.window,
                (200, 200, 200),
                pygame.Rect(self.player_pos["x"], self.player_pos["y"], 50, 50),
            )
            pygame.display.flip()
