#!/usr/bin/python3

import pygame
import sys

from src.block import Block
from src.constants import (
    W,
    H,
    GAME_TITLE,
    FPS,
    BLOCK_SIDE,
    PLACEHOLDER_COLOR,
    PLACEHOLDER_COLOR_2,
)
from src.player import Player


def init_pygame(name, w, h, debug):
    pygame.display.init()
    pygame.font.init()
    pygame.display.set_caption(name)
    pygame.mouse.set_visible(False)
    window = None
    if debug:
        window = pygame.display.set_mode((w, h), vsync=1)
    else:
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


def create_ground_blocks():
    nb_blocks_in_width = int(W // BLOCK_SIDE) + 1
    ground_height = H - BLOCK_SIDE
    ground_blocks = [
        Block(
            i * BLOCK_SIDE,
            ground_height,
        )
        for i in range(nb_blocks_in_width)
    ]
    return ground_blocks


def blocks_from_map(m):
    blocks = []
    for i in range(len(m)):
        for j in range(len(m[i])):
            if m[i][j] == 1:
                x = j * BLOCK_SIDE
                y = i * BLOCK_SIDE
                blocks.append(Block(x, y))
    return blocks


def print_map(m):
    for line in m:
        print(line)


def main():
    # Command line arguments
    DEBUG = False
    if len(sys.argv) > 1 and sys.argv[1] == "debug":
        DEBUG = True

    # Init and create objects
    window, clock = init_pygame(GAME_TITLE, W, H, DEBUG)
    player = Player()

    # Map
    blocks_map = [[0 for i in range(16)] for _ in range(9)]
    blocks_map[8] = [1 for i in range(16)]
    blocks_map[7][0] = 1
    blocks_map[7][1] = 1
    blocks_map[5][3] = 1
    blocks_map[5][4] = 1
    blocks_map[5][5] = 1
    blocks_map[5][6] = 1
    blocks = blocks_from_map(blocks_map)

    # Add sprites to groups
    blocks_group = pygame.sprite.Group()
    player_group = pygame.sprite.Group()

    player_group.add(player)
    blocks_group.add(blocks)

    # Load background
    background = pygame.image.load("assets/background.png")
    background = pygame.transform.scale(background, [W, H])

    # Game loop
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
        window.blit(background, (0, 0))
        blocks_group.draw(window)
        player_group.draw(window)
        if DEBUG:
            pygame.draw.rect(window, PLACEHOLDER_COLOR, player.hitrect, 1)
            pygame.draw.rect(window, PLACEHOLDER_COLOR_2, player.rect, 1)
        pygame.display.flip()


if __name__ == "__main__":
    main()
