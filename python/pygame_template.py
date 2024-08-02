#!/usr/bin/python3

import pygame

from pygame.locals import QUIT


W = 1000
H = 800
FPS = 60.0


def main():
    # Init
    pygame.display.init()
    pygame.display.set_caption("PyGame basic template")
    window = pygame.display.set_mode((W, H))
    clock = pygame.time.Clock()

    # Main loop
    while True:
        # Events
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            break

        # Keyboard events
        keys = pygame.key.get_pressed()
        if keys[pygame.K_ESCAPE]:
            break

        # Update
        dt = clock.tick(FPS)

        # Display
        window.fill((0, 0, 0))
        pygame.display.flip()


if __name__ == "__main__":
    main()
