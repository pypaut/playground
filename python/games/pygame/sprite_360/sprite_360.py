#!/usr/bin/python3

import math as m
import numpy as np
import os
import pygame

from pygame.locals import QUIT, MOUSEBUTTONDOWN


W = 1000
H = 800
FPS = 60.0
NB_FRAMES = 5
CHAR_POS = np.array([H // 2, W // 2])
DIR_LIMITS = [
    0,
    m.pi / 8,
    m.pi * 3 / 8,
    m.pi * 5 / 8,
    m.pi * 7 / 8,
    m.pi,
]
EPSILON = 0.1


def load_sprites(path):
    """
    Load sprites as PyGame images in sorted list, from @path
    """
    filenames = sorted(os.listdir(path))
    sprites = [pygame.image.load(os.path.join(path, f)) for f in filenames]
    return sprites


def oppose_dir(direction):
    """
    Return opposite direction
    """
    if direction == "up_right":
        return "down_right"
    elif direction == "up":
        return "down"
    elif direction == "up_left":
        return "down_left"
    else:
        return direction


def main():
    # Init
    pygame.display.init()
    pygame.display.set_caption("PyGame sprites animations")
    window = pygame.display.set_mode((W, H))
    clock = pygame.time.Clock()
    is_pressed = False

    # Load sprites
    walk_files = os.listdir("sprites/walk")
    sprites_walk = {f: load_sprites(f"sprites/walk/{f}") for f in walk_files}
    idle_files = os.listdir("sprites/idle")
    sprites_idle = {
        f[:-4]: [pygame.image.load(os.path.join("sprites/idle", f))]
        for f in idle_files
    }
    sprites = {
        "walk": sprites_walk,
        "idle": sprites_idle,
    }

    # Variables for main loop
    direction = "down_right"
    animation = "idle"
    frame_counter = 0  # Count frames for each sprite
    index_sprite = 0  # Current sprite index to display

    # Main loop
    while True:
        # Events
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            break

        # Mouse event
        left_mouse = pygame.mouse.get_pressed(num_buttons=3)[0]
        if left_mouse:
            animation = "walk"
            # Compute direction
            mouse_pos = np.array(pygame.mouse.get_pos())
            direction_vec = mouse_pos - CHAR_POS
            direction_vec = direction_vec / np.linalg.norm(direction_vec, 2)
            complex_pos = complex(direction_vec[0], direction_vec[1])
            angle = -np.angle([complex_pos])[0]

            if DIR_LIMITS[0] < abs(angle) < DIR_LIMITS[1]:
                direction = "right"
            if DIR_LIMITS[1] < abs(angle) < DIR_LIMITS[2]:
                direction = "up_right"
            if DIR_LIMITS[2] < abs(angle) < DIR_LIMITS[3]:
                direction = "up"
            if DIR_LIMITS[3] < abs(angle) < DIR_LIMITS[4]:
                direction = "up_left"
            if DIR_LIMITS[4] < abs(angle) < DIR_LIMITS[5]:
                direction = "left"
            if angle < 0:
                direction = oppose_dir(direction)
        else:
            animation = "idle"

        # Update sprite index
        dt = clock.tick(FPS)
        if frame_counter >= NB_FRAMES:
            index_sprite += 1
            index_sprite %= len(sprites[animation][direction])
            frame_counter = 0
        else:
            frame_counter += 1

        if animation == "idle":
            index_sprite = 0

        # Display
        window.fill((0, 0, 0))
        window.blit(
            sprites[animation][direction][index_sprite],
            CHAR_POS,
        )
        pygame.display.flip()


if __name__ == "__main__":
    main()
