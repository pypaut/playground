#!/usr/bin/python3

import pygame

from src.blocks import Blocks
from src.constants import W, H, GAME_TITLE, FPS
from src.player import Player


def init_pygame(name, w, h):
    pygame.display.init()
    pygame.font.init()
    pygame.display.set_caption(name)
    pygame.mouse.set_visible(False)
    window = pygame.display.set_mode((w, h), flags=pygame.SCALED, vsync=1)
    clock = pygame.time.Clock()
    return window, clock


def check_quit_event(keys, events):
    # Check quit event
    if pygame.QUIT in events:
        return True

    # Check ESC
    if keys[pygame.K_ESCAPE]:
        return True

    return False


def main():
    window, clock = init_pygame(GAME_TITLE, W, H)
    player = Player(W, H)
    blocks = Blocks()

    while True:
        dt = clock.tick(FPS)

        # Events
        events = pygame.event.get()
        events_types = [e.type for e in events]
        keys = pygame.key.get_pressed()
        if check_quit_event(keys, events_types):
            break

        player.events(keys)

        # Update
        player.update(keys, dt, blocks)

        # Draw
        window.fill((0, 0, 0))
        blocks.draw(window)
        player.draw(window)
        pygame.display.flip()


if __name__ == "__main__":
    main()
