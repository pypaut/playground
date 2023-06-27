#!/usr/bin/python3

import pygame

from src.block import Block
from src.constants import W, H, GAME_TITLE, FPS, BLOCK_SIDE
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


def main():
    # Init and create objects
    window, clock = init_pygame(GAME_TITLE, W, H)
    player = Player()
    blocks = []

    ground_blocks = create_ground_blocks()
    blocks += ground_blocks

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
        pygame.display.flip()


if __name__ == "__main__":
    main()
