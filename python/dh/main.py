#!/usr/bin/python3

import pygame

from src.game import game
from src.menu import menu
from src.constants import W, H


def init_pygame(name, w, h):
    pygame.display.init()
    pygame.font.init()
    pygame.display.set_caption(name)
    window = pygame.display.set_mode((w, h))
    clock = pygame.time.Clock()
    return window, clock


def main():
    window, clock = init_pygame("Dark Hunter", W, H)

    while True:
        command = menu(
            window, clock, "DARK HUNTER", [("PLAY", "play"), ("QUIT", "quit")]
        )

        if command == "play":
            game(window, clock)
        else:  # command == "quit"
            return


if __name__ == "__main__":
    main()
