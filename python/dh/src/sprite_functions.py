import os
import pygame


def left_sprites(sprites):
    """
    Return new list containing vertically split sprites from @sprites
    """
    return [pygame.transform.flip(i.copy(), True, False) for i in sprites]


def load_sprites(path):
    """
    Load sprites as PyGame images in sorted list, from @path
    """
    filenames = sorted(os.listdir(path))
    sprites = [
        pygame.image.load(os.path.join(path, f)).convert() for f in filenames
    ]
    return sprites
