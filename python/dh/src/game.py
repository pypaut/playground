import pygame

from src.block import Block
from src.constants import FPS, W, H
from src.menu import pause_menu
from src.player import Player


def game(window, clock):
    player = Player()
    ground = Block(0, H * 3 / 4, W, 50)
    block_1 = Block(W * 2 / 3, ground.rect.y - 50, 50, 50)
    block_2 = Block(W * 1 / 3, ground.rect.y - 50, 50, 50)
    blocks = [ground, block_1, block_2]

    # Main loop
    while True:
        """
        Events
        """
        # Quit
        events = pygame.event.get()
        events_types = [e.type for e in events]
        if pygame.QUIT in events_types:
            return  # TODO : Are you sure?

        # Events
        keys = pygame.key.get_pressed()
        if keys[pygame.K_ESCAPE]:
            command = pause_menu(window, clock)
            if command == "quit":
                return
            continue

        # Update
        dt = clock.tick(FPS)
        player.update(keys, blocks, dt)

        # Display
        window.fill((0, 0, 0))
        window.blit(
            player.get_sprite(),
            (player.rect.x, player.rect.y),
        )
        for b in blocks:
            b.draw(window)
        pygame.display.flip()
