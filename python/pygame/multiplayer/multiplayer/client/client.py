import configparser
import json
import math
import pygame
import socket

from pygame.locals import QUIT


class Client:
    def __init__(self):
        # Read config
        config = configparser.ConfigParser()
        with open("client.ini", "r") as configfile:
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

        # Game logic
        self.players = {}

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

            player_dir = self.get_player_dir(keys)
            self.send_direction(player_dir)
            self.receive_positions()
            self.draw()

    def get_player_dir(self, keys):
        player_dir = {"x": 0.0, "y": 0.0}

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

        return player_dir

    def send_direction(self, player_dir):
        self.socket.send(f"{json.dumps(player_dir)}".encode("utf-8"))

    def receive_positions(self):
        server_msg = self.socket.recv(1024)
        if server_msg != b"":
            self.players = json.loads(server_msg)

    def draw(self):
        self.window.fill((0, 0, 0))
        for _, p in self.players.items():
            pygame.draw.rect(
                self.window,
                (200, 200, 200),
                pygame.Rect(p["x"], p["y"], 50, 50),
            )
        pygame.display.flip()
