#!/usr/bin/python3

import os
import pygame

from pygame.locals import QUIT


W = 1000
H = 800
FPS = 60.0
NB_FRAMES = 5


def load_sprites(path):
    """
    Load sprites as PyGame images in sorted list, from @path
    """
    filenames = sorted(os.listdir(path))
    sprites = [pygame.image.load(os.path.join(path, f)) for f in filenames]
    return sprites


def left_sprites(sprites):
    """
    Return new list containing vertically split sprites from @sprites
    """
    return [pygame.transform.flip(i.copy(), True, False) for i in sprites]


def main():
    # Init
    pygame.display.init()
    pygame.display.set_caption("PyGame sprites animations")
    window = pygame.display.set_mode((W, H))
    clock = pygame.time.Clock()

    # Load sprites
    animations_files = os.listdir("sprites")
    sprites = {
        "right": {f: load_sprites(f"sprites/{f}") for f in animations_files},
    }
    sprites["left"] = {
        a: left_sprites(s) for (a, s) in sprites["right"].items()
    }

    # Variables for main loop
    direction = "right"
    animation = "idle"
    is_attacking = False
    attack_counter = 0
    index_sprite = 0
    frame_counter = 0
    nb_sprites = len(sprites[direction][animation])

    # Main loop
    while True:
        # Events
        events = pygame.event.get()
        if QUIT in [e.type for e in events]:
            break

        # Keyboard event
        keys = pygame.key.get_pressed()
        if keys[pygame.K_a] or keys[pygame.K_LEFT]:
            direction = "left"
            if keys[pygame.K_LSHIFT]:
                animation = "run"
            else:
                animation = "walk"
        elif keys[pygame.K_d] or keys[pygame.K_RIGHT]:
            direction = "right"
            if keys[pygame.K_LSHIFT]:
                animation = "run"
            else:
                animation = "walk"
        else:
            if not True in keys or keys[pygame.K_LSHIFT]:
                animation = "idle"

        if keys[pygame.K_e] and not is_attacking:
            is_attacking = True
            animation = "attack"
            frame_counter = 0
            index_sprite = 0

        if is_attacking:
            animation = "attack"

        # Update sprite index
        dt = clock.tick(FPS)
        if frame_counter >= NB_FRAMES:
            index_sprite += 1
            # End of attack
            if animation == "attack" and index_sprite == nb_sprites:
                animation = "idle"
                is_attacking = False
            frame_counter = 0
            index_sprite %= nb_sprites
        else:
            frame_counter += 1

        # Display
        window.fill((0, 0, 0))
        window.blit(
            sprites[direction][animation][index_sprite],
            [H // 2, W // 2],
        )
        pygame.display.flip()


if __name__ == "__main__":
    main()
